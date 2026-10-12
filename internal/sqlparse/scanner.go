package sqlparse

import "strings"

// FindMatchingParen returns the closing parenthesis paired with open. SQL
// strings, bracketed identifiers, line comments, and nested block comments do
// not affect the nesting depth.
func FindMatchingParen(input string, open int) (int, bool) {
	if open < 0 || open >= len(input) || input[open] != '(' {
		return -1, false
	}
	depth := 0
	state := scanNormal
	blockDepth := 0
	for i := open; i < len(input); i++ {
		ch := input[i]
		next := byte(0)
		if i+1 < len(input) {
			next = input[i+1]
		}
		switch state {
		case scanString:
			if ch == '\'' {
				if next == '\'' {
					i++
					continue
				}
				state = scanNormal
			}
			continue
		case scanBracket:
			if ch == ']' {
				if next == ']' {
					i++
					continue
				}
				state = scanNormal
			}
			continue
		case scanLineComment:
			if ch == '\n' {
				state = scanNormal
			}
			continue
		case scanBlockComment:
			if ch == '/' && next == '*' {
				blockDepth++
				i++
				continue
			}
			if ch == '*' && next == '/' {
				blockDepth--
				i++
				if blockDepth == 0 {
					state = scanNormal
				}
			}
			continue
		}

		switch {
		case ch == '\'':
			state = scanString
		case ch == '[':
			state = scanBracket
		case ch == '-' && next == '-':
			state = scanLineComment
			i++
		case ch == '/' && next == '*':
			state = scanBlockComment
			blockDepth = 1
			i++
		case ch == '(':
			depth++
		case ch == ')':
			depth--
			if depth == 0 {
				return i, true
			}
		}
	}
	return -1, false
}

// SplitTopLevelComma splits a SQL list only on commas outside nested
// expressions, literals, comments, and bracketed identifiers.
func SplitTopLevelComma(input string) []string {
	parts := make([]string, 0)
	start := 0
	depth := 0
	state := scanNormal
	blockDepth := 0
	for i := 0; i < len(input); i++ {
		ch := input[i]
		next := byte(0)
		if i+1 < len(input) {
			next = input[i+1]
		}
		switch state {
		case scanString:
			if ch == '\'' {
				if next == '\'' {
					i++
					continue
				}
				state = scanNormal
			}
			continue
		case scanBracket:
			if ch == ']' {
				if next == ']' {
					i++
					continue
				}
				state = scanNormal
			}
			continue
		case scanLineComment:
			if ch == '\n' {
				state = scanNormal
			}
			continue
		case scanBlockComment:
			if ch == '/' && next == '*' {
				blockDepth++
				i++
				continue
			}
			if ch == '*' && next == '/' {
				blockDepth--
				i++
				if blockDepth == 0 {
					state = scanNormal
				}
			}
			continue
		}

		switch {
		case ch == '\'':
			state = scanString
		case ch == '[':
			state = scanBracket
		case ch == '-' && next == '-':
			state = scanLineComment
			i++
		case ch == '/' && next == '*':
			state = scanBlockComment
			blockDepth = 1
			i++
		case ch == '(':
			depth++
		case ch == ')' && depth > 0:
			depth--
		case ch == ',' && depth == 0:
			parts = append(parts, input[start:i])
			start = i + 1
		}
	}
	parts = append(parts, input[start:])
	return parts
}

// TopLevelLineOffsets returns byte offsets of non-first physical lines that
// begin while the scanner is outside parentheses, literals, identifiers, and
// block comments. It is intended for conservative statement-list recovery.
func TopLevelLineOffsets(input string) []int {
	offsets := make([]int, 0)
	state := scanNormal
	depth := 0
	blockDepth := 0
	for i := 0; i < len(input); i++ {
		ch := input[i]
		switch state {
		case scanNormal:
			switch {
			case ch == '\'':
				state = scanString
			case ch == '[':
				state = scanBracket
			case ch == '-' && i+1 < len(input) && input[i+1] == '-':
				state = scanLineComment
				i++
			case ch == '/' && i+1 < len(input) && input[i+1] == '*':
				state = scanBlockComment
				blockDepth = 1
				i++
			case ch == '(':
				depth++
			case ch == ')' && depth > 0:
				depth--
			case (ch == '\n' || ch == '\r') && depth == 0:
				if ch == '\r' && i+1 < len(input) && input[i+1] == '\n' {
					i++
				}
				if i+1 < len(input) {
					offsets = append(offsets, i+1)
				}
			}
		case scanString:
			if ch == '\'' {
				if i+1 < len(input) && input[i+1] == '\'' {
					i++
				} else {
					state = scanNormal
				}
			}
		case scanBracket:
			if ch == ']' {
				if i+1 < len(input) && input[i+1] == ']' {
					i++
				} else {
					state = scanNormal
				}
			}
		case scanLineComment:
			if ch == '\n' || ch == '\r' {
				state = scanNormal
				if ch == '\r' && i+1 < len(input) && input[i+1] == '\n' {
					i++
				}
				if depth == 0 && i+1 < len(input) {
					offsets = append(offsets, i+1)
				}
			}
		case scanBlockComment:
			if ch == '/' && i+1 < len(input) && input[i+1] == '*' {
				blockDepth++
				i++
			} else if ch == '*' && i+1 < len(input) && input[i+1] == '/' {
				blockDepth--
				i++
				if blockDepth == 0 {
					state = scanNormal
				}
			}
		}
	}
	return offsets
}

// RewriteIdentifierExact rewrites one exact-case SQL identifier outside
// literals and comments. Bracket escaping is preserved for the replacement.
func RewriteIdentifierExact(input, oldName, newName string) string {
	if oldName == "" || oldName == newName {
		return input
	}
	var out strings.Builder
	out.Grow(len(input))
	state := scanNormal
	blockDepth := 0
	for i := 0; i < len(input); {
		ch := input[i]
		next := byte(0)
		if i+1 < len(input) {
			next = input[i+1]
		}
		switch state {
		case scanString:
			out.WriteByte(ch)
			i++
			if ch == '\'' {
				if next == '\'' {
					out.WriteByte(next)
					i++
				} else {
					state = scanNormal
				}
			}
			continue
		case scanLineComment:
			out.WriteByte(ch)
			i++
			if ch == '\n' {
				state = scanNormal
			}
			continue
		case scanBlockComment:
			out.WriteByte(ch)
			i++
			if ch == '/' && next == '*' {
				out.WriteByte(next)
				i++
				blockDepth++
			} else if ch == '*' && next == '/' {
				out.WriteByte(next)
				i++
				blockDepth--
				if blockDepth == 0 {
					state = scanNormal
				}
			}
			continue
		}

		if ch == '\'' {
			state = scanString
			out.WriteByte(ch)
			i++
			continue
		}
		if ch == '-' && next == '-' {
			state = scanLineComment
			out.WriteString("--")
			i += 2
			continue
		}
		if ch == '/' && next == '*' {
			state = scanBlockComment
			blockDepth = 1
			out.WriteString("/*")
			i += 2
			continue
		}
		if ch == '[' {
			end := i + 1
			var name strings.Builder
			for end < len(input) {
				if input[end] == ']' {
					if end+1 < len(input) && input[end+1] == ']' {
						name.WriteByte(']')
						end += 2
						continue
					}
					end++
					break
				}
				name.WriteByte(input[end])
				end++
			}
			if name.String() == oldName {
				out.WriteByte('[')
				out.WriteString(strings.ReplaceAll(newName, "]", "]]"))
				out.WriteByte(']')
			} else {
				out.WriteString(input[i:end])
			}
			i = end
			continue
		}
		if isIdentifierByte(ch) {
			end := i + 1
			for end < len(input) && isIdentifierByte(input[end]) {
				end++
			}
			token := input[i:end]
			if token == oldName {
				out.WriteString(newName)
			} else {
				out.WriteString(token)
			}
			i = end
			continue
		}
		out.WriteByte(ch)
		i++
	}
	return out.String()
}

// ContainsIdentifierExact reports whether input contains the exact identifier
// token outside literals and comments.
func ContainsIdentifierExact(input, name string) bool {
	if name == "" {
		return false
	}
	return RewriteIdentifierExact(input, name, name+"\x00") != input
}

type scanState uint8

const (
	scanNormal scanState = iota
	scanString
	scanBracket
	scanLineComment
	scanBlockComment
)

func isIdentifierByte(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_' || ch == '#' || ch == '@' || ch == '$'
}
