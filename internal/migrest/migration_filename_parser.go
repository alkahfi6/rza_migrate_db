package migrest

import (
	"strconv"
	"strings"
	"unicode"
)

type migrationFileMeta struct {
	FileName  string
	Version   string
	Category  string
	Object    string
	Direction string
	VK        int
	IsCore    bool
}

type parseIssue struct {
	Rule    string
	Message string
}

func parseMigrationFileMeta(name string) (migrationFileMeta, *parseIssue) {
	if name == coreAuditUpName {
		return migrationFileMeta{
			FileName:  name,
			Version:   "000000",
			Category:  "core",
			Direction: "up",
			VK:        0,
			IsCore:    true,
		}, nil
	}
	if name == coreAuditDownName {
		return migrationFileMeta{
			FileName:  name,
			Version:   "000000",
			Category:  "core",
			Direction: "down",
			VK:        0,
			IsCore:    true,
		}, nil
	}

	direction := ""
	base := ""
	lowerName := strings.ToLower(name)
	switch {
	case strings.HasSuffix(lowerName, ".up.sql"):
		direction = "up"
		base = name[:len(name)-len(".up.sql")]
	case strings.HasSuffix(lowerName, ".down.sql"):
		direction = "down"
		base = name[:len(name)-len(".down.sql")]
	default:
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_MIGRATION_FILENAME",
			Message: "migration filename must end with .up.sql or .down.sql",
		}
	}

	parts := strings.Split(base, "_")
	if len(parts) < 4 {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_MIGRATION_FILENAME",
			Message: "migration filename has invalid token count",
		}
	}

	version := parts[0]
	if len(version) != 17 || !isAllDigits(version) {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_VERSION_FORMAT",
			Message: "version must be 17 digits for generated migration files",
		}
	}

	category := parts[1]
	if _, ok := migrationAllowedCategories()[category]; !ok {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_CATEGORY",
			Message: "migration category is not in whitelist",
		}
	}

	vkToken := parts[len(parts)-1]
	if len(vkToken) < 2 || vkToken[0] != 'v' {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_VK_TOKEN",
			Message: "migration filename must end with v{k}",
		}
	}
	vkValue, err := strconv.Atoi(vkToken[1:])
	if err != nil || vkValue < 1 {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_VK_TOKEN",
			Message: "migration vk must be integer >= 1",
		}
	}

	nameToken := strings.Join(parts[2:len(parts)-1], "_")
	if strings.TrimSpace(nameToken) == "" {
		return migrationFileMeta{}, &parseIssue{
			Rule:    "INVALID_OBJECT_TOKEN",
			Message: "migration object token is empty",
		}
	}

	if category == "patch" {
		if !isValidPatchToken(nameToken) {
			return migrationFileMeta{}, &parseIssue{
				Rule:    "INVALID_PATCH_NAME_TOKEN",
				Message: "patch token contains invalid character",
			}
		}
	} else {
		if !isValidObjectToken(nameToken, migrationAllowedObjectTokenRunes()) {
			return migrationFileMeta{}, &parseIssue{
				Rule:    "INVALID_OBJECT_TOKEN",
				Message: "object token contains invalid character",
			}
		}
		if !isValidDBOMigrationObjectToken(nameToken) {
			return migrationFileMeta{}, &parseIssue{
				Rule:    "MIGRATION_OBJECT_SCHEMA_NOT_DBO",
				Message: "migration object token must be dbo.<ObjectName>",
			}
		}
	}

	return migrationFileMeta{
		FileName:  name,
		Version:   version,
		Category:  category,
		Object:    nameToken,
		Direction: direction,
		VK:        vkValue,
		IsCore:    false,
	}, nil
}

func isAllDigits(input string) bool {
	if input == "" {
		return false
	}
	for _, ch := range input {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func isValidPatchToken(token string) bool {
	for _, ch := range token {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '-' {
			continue
		}
		return false
	}
	return strings.TrimSpace(token) != ""
}

func isValidObjectToken(token string, allowed map[rune]struct{}) bool {
	if strings.TrimSpace(token) == "" {
		return false
	}
	for _, ch := range token {
		if _, ok := allowed[ch]; !ok {
			return false
		}
	}
	return true
}

func isValidDBOMigrationObjectToken(token string) bool {
	trimmed := strings.TrimSpace(token)
	if trimmed == "" {
		return false
	}
	parts := strings.Split(trimmed, ".")
	if len(parts) != 2 {
		return false
	}
	if parts[0] != "dbo" {
		return false
	}
	return strings.TrimSpace(parts[1]) != ""
}
