package getobject

import (
	"fmt"
	"regexp"
	"strings"
)

type moduleHeaderRule struct {
	canonical string
	re        *regexp.Regexp
}

var moduleHeaderRules = map[string]moduleHeaderRule{
	"VIEW": {
		canonical: "CREATE OR ALTER VIEW",
		re:        regexp.MustCompile(`(?is)^(CREATE|ALTER)\s+(OR\s+ALTER\s+)?VIEW\b`),
	},
	"FUNCTION": {
		canonical: "CREATE OR ALTER FUNCTION",
		re:        regexp.MustCompile(`(?is)^(CREATE|ALTER)\s+(OR\s+ALTER\s+)?FUNCTION\b`),
	},
	"PROCEDURE": {
		canonical: "CREATE OR ALTER PROCEDURE",
		re:        regexp.MustCompile(`(?is)^(CREATE|ALTER)\s+(OR\s+ALTER\s+)?PROC(?:EDURE)?\b`),
	},
}

var moduleSetDirectiveRe = regexp.MustCompile(`(?is)^SET\s+[A-Z_][A-Z0-9_]*\s+(ON|OFF)\s*;?`)
var goLineRe = regexp.MustCompile(`(?is)^GO(?:\s+--.*)?$`)

func ensureCreateOrAlter(definition, kind string) (string, error) {
	kindUpper := strings.ToUpper(strings.TrimSpace(kind))
	if kindUpper == "" {
		return "", fmt.Errorf("kind is required")
	}

	rule, ok := moduleHeaderRules[kindUpper]
	if !ok {
		return "", fmt.Errorf("unsupported kind %s", kindUpper)
	}

	normalized := normalizeSQL(definition)
	normalized = strings.TrimPrefix(normalized, "\uFEFF")
	trimmed := trimLeadingModulePreamble(normalized)

	loc := rule.re.FindStringIndex(trimmed)
	if loc == nil || loc[0] != 0 {
		return "", fmt.Errorf("definition must start with CREATE/ALTER %s", kindUpper)
	}

	rest := strings.TrimLeft(trimmed[loc[1]:], " \t\n")
	objectToken, consumed := consumeLeadingModuleObjectToken(rest)
	if strings.TrimSpace(objectToken) == "" {
		return "", fmt.Errorf("definition must include object name after %s", rule.canonical)
	}
	objectToken = canonicalizeModuleObjectToken(objectToken)
	if objectToken == "" {
		return "", fmt.Errorf("definition object name is invalid")
	}
	remaining := strings.TrimLeft(rest[consumed:], " \t\n")
	rewritten := rule.canonical
	rewritten += " " + objectToken
	if remaining != "" {
		rewritten += " " + remaining
	}
	return normalizeSQL(rewritten), nil
}

// EnsureCreateOrAlter normalizes a SQL Server module header to CREATE OR ALTER.
func EnsureCreateOrAlter(definition, kind string) (string, error) {
	return ensureCreateOrAlter(definition, kind)
}

func trimLeadingModulePreamble(input string) string {
	s := input
	for {
		s = strings.TrimLeft(s, " \t\n")
		if s == "" {
			return s
		}

		advanced := false

		if strings.HasPrefix(s, "--") {
			if pos := strings.IndexByte(s, '\n'); pos >= 0 {
				s = s[pos+1:]
			} else {
				s = ""
			}
			advanced = true
		} else if strings.HasPrefix(s, "/*") {
			if pos := strings.Index(s, "*/"); pos >= 0 {
				s = s[pos+2:]
				advanced = true
			} else {
				return s
			}
		} else if loc := moduleSetDirectiveRe.FindStringIndex(s); loc != nil && loc[0] == 0 {
			s = s[loc[1]:]
			advanced = true
		} else if rest, ok := consumeLeadingGoLine(s); ok {
			s = rest
			advanced = true
		}

		if !advanced {
			return s
		}
	}
}

func consumeLeadingGoLine(input string) (string, bool) {
	lineEnd := strings.IndexByte(input, '\n')
	line := input
	rest := ""
	if lineEnd >= 0 {
		line = input[:lineEnd]
		rest = input[lineEnd+1:]
	}
	if goLineRe.MatchString(strings.TrimSpace(line)) {
		return rest, true
	}
	return input, false
}

func consumeLeadingModuleObjectToken(input string) (string, int) {
	if input == "" {
		return "", 0
	}
	i := 0
	for i < len(input) {
		switch input[i] {
		case ' ', '\t', '\n', '\r', '(', ',', ';':
			return input[:i], i
		default:
			i++
		}
	}
	return input, len(input)
}

func canonicalizeModuleObjectToken(token string) string {
	trimmed := strings.TrimSpace(strings.TrimRight(token, ",;"))
	if trimmed == "" {
		return ""
	}
	withoutBrackets := strings.ReplaceAll(strings.ReplaceAll(trimmed, "[", ""), "]", "")
	withoutBrackets = strings.ReplaceAll(withoutBrackets, " ", "")
	if withoutBrackets == "" {
		return ""
	}

	partsRaw := strings.Split(withoutBrackets, ".")
	parts := make([]string, 0, len(partsRaw))
	for _, part := range partsRaw {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		parts = append(parts, part)
	}
	if len(parts) == 0 {
		return ""
	}

	schema := "dbo"
	object := ""
	switch len(parts) {
	case 1:
		object = parts[0]
	case 2:
		schema = parts[0]
		object = parts[1]
	default:
		schema = parts[len(parts)-2]
		object = parts[len(parts)-1]
	}
	if schema != "dbo" {
		return ""
	}
	object = strings.TrimSpace(object)
	if object == "" || strings.Contains(object, ".") {
		return ""
	}
	return "dbo." + object
}
