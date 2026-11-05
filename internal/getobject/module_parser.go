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
	_, end, ok := parseCatalogModuleName(input)
	if !ok {
		return "", 0
	}
	return input[:end], end
}

func canonicalizeModuleObjectToken(token string) string {
	parts, _, ok := parseCatalogModuleName(token)
	if !ok || len(parts) < 1 || len(parts) > 2 {
		return ""
	}
	schema, object := "dbo", parts[len(parts)-1]
	if len(parts) == 2 {
		schema = parts[0]
	}
	if schema != "dbo" {
		return ""
	}
	return qualifiedName(schema, object)
}

// Read identifier parts without treating punctuation inside delimiters as syntax.
func parseCatalogModuleName(input string) ([]string, int, bool) {
	var parts []string
	i := 0
	for {
		for i < len(input) && strings.ContainsRune(" \t\r\n", rune(input[i])) {
			i++
		}
		if i >= len(input) {
			return nil, 0, false
		}
		var part strings.Builder
		if input[i] == '[' || input[i] == '"' {
			close := byte(']')
			if input[i] == '"' {
				close = '"'
			}
			i++
			closed := false
			for i < len(input) {
				b := input[i]
				i++
				if b == close {
					if i < len(input) && input[i] == close {
						part.WriteByte(close)
						i++
						continue
					}
					closed = true
					break
				}
				part.WriteByte(b)
			}
			if !closed {
				return nil, 0, false
			}
		} else {
			start := i
			for i < len(input) && !strings.ContainsRune(". \t\r\n(,;", rune(input[i])) {
				i++
			}
			part.WriteString(input[start:i])
		}
		if part.Len() == 0 {
			return nil, 0, false
		}
		parts = append(parts, part.String())
		end := i
		for i < len(input) && strings.ContainsRune(" \t\r\n", rune(input[i])) {
			i++
		}
		if i >= len(input) || input[i] != '.' {
			return parts, end, true
		}
		i++
	}
}
