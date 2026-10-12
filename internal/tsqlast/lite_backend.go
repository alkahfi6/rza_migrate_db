package tsqlast

import (
	"strings"

	"rza_migrate_db/internal/sqlfmt"
)

func formatWithLite(input string, cfg FormatConfig) string {
	out := normalizeSQL(input)
	if cfg.StripComments {
		out = StripSQLComments(out)
	}
	if cfg.Reindent {
		out = ReindentTSQL(out, cfg.IndentSize)
	}
	return normalizeSQL(out)
}

func normalizeSQL(input string) string {
	return sqlfmt.NormalizeSQLText(input, true, true)
}

func StripSQLComments(input string) string {
	s := normalizeSQL(input)
	var b strings.Builder
	b.Grow(len(s))

	inLineComment := false
	blockDepth := 0
	inString := false
	inBracket := false

	for i := 0; i < len(s); {
		ch := s[i]
		next := byte(0)
		if i+1 < len(s) {
			next = s[i+1]
		}

		if blockDepth > 0 {
			if ch == '/' && next == '*' {
				blockDepth++
				i += 2
				continue
			}
			if ch == '*' && next == '/' {
				blockDepth--
				i += 2
				continue
			}
			if ch == '\n' {
				b.WriteByte('\n')
			}
			i++
			continue
		}

		if inLineComment {
			if ch == '\n' {
				inLineComment = false
				b.WriteByte('\n')
			}
			i++
			continue
		}

		if inString {
			b.WriteByte(ch)
			if ch == '\'' {
				if next == '\'' {
					b.WriteByte(next)
					i += 2
					continue
				}
				inString = false
			}
			i++
			continue
		}

		if inBracket {
			b.WriteByte(ch)
			if ch == ']' {
				if next == ']' {
					b.WriteByte(next)
					i += 2
					continue
				}
				inBracket = false
			}
			i++
			continue
		}

		if ch == '-' && next == '-' {
			inLineComment = true
			i += 2
			continue
		}
		if ch == '/' && next == '*' {
			blockDepth = 1
			i += 2
			continue
		}
		if ch == '\'' {
			inString = true
			b.WriteByte(ch)
			i++
			continue
		}
		if ch == '[' {
			inBracket = true
			b.WriteByte(ch)
			i++
			continue
		}

		b.WriteByte(ch)
		i++
	}

	return normalizeSQL(b.String())
}

type contextTokenType int

const (
	contextBlock contextTokenType = iota
	contextCase
)

type wordToken struct {
	Text  string
	Start int
	End   int
}

type protectedLine struct {
	Raw             string
	Code            string
	Preserve        bool
	ContainsComment bool
}

// analyzeProtectedLines masks SQL trivia for structural analysis while
// retaining the exact original line for output. State is deliberately carried
// across lines so nested block comments and multiline literals cannot leak
// BEGIN/END/CASE tokens into the indentation stack.
func analyzeProtectedLines(lines []string) []protectedLine {
	result := make([]protectedLine, len(lines))
	blockDepth := 0
	inString := false
	inBracket := false
	inQuoted := false
	for lineIndex, raw := range lines {
		masked := []byte(raw)
		startedInMultilineToken := blockDepth > 0 || inString || inBracket || inQuoted
		containsComment := blockDepth > 0
		for i := 0; i < len(raw); {
			ch := raw[i]
			next := byte(0)
			if i+1 < len(raw) {
				next = raw[i+1]
			}

			if blockDepth > 0 {
				masked[i] = ' '
				if ch == '/' && next == '*' {
					masked[i+1] = ' '
					blockDepth++
					i += 2
					continue
				}
				if ch == '*' && next == '/' {
					masked[i+1] = ' '
					blockDepth--
					i += 2
					continue
				}
				i++
				continue
			}
			if inString {
				masked[i] = ' '
				if ch == '\'' {
					if next == '\'' {
						masked[i+1] = ' '
						i += 2
						continue
					}
					inString = false
				}
				i++
				continue
			}
			if inBracket {
				masked[i] = ' '
				if ch == ']' {
					if next == ']' {
						masked[i+1] = ' '
						i += 2
						continue
					}
					inBracket = false
				}
				i++
				continue
			}
			if inQuoted {
				masked[i] = ' '
				if ch == '"' {
					if next == '"' {
						masked[i+1] = ' '
						i += 2
						continue
					}
					inQuoted = false
				}
				i++
				continue
			}

			if ch == '-' && next == '-' {
				containsComment = true
				for j := i; j < len(masked); j++ {
					masked[j] = ' '
				}
				break
			}
			if ch == '/' && next == '*' {
				containsComment = true
				masked[i], masked[i+1] = ' ', ' '
				blockDepth = 1
				i += 2
				continue
			}
			if ch == '\'' {
				masked[i] = ' '
				inString = true
				i++
				continue
			}
			if ch == '[' {
				masked[i] = ' '
				inBracket = true
				i++
				continue
			}
			if ch == '"' {
				masked[i] = ' '
				inQuoted = true
				i++
				continue
			}
			i++
		}
		continuesMultilineToken := blockDepth > 0 || inString || inBracket || inQuoted
		result[lineIndex] = protectedLine{
			Raw:             raw,
			Code:            string(masked),
			Preserve:        containsComment || startedInMultilineToken || continuesMultilineToken,
			ContainsComment: containsComment,
		}
	}
	return result
}

// ContainsSQLComment reports lexical SQL comments without mistaking markers
// inside string literals or bracketed/quoted identifiers for comments.
func ContainsSQLComment(input string) bool {
	lines := splitSQLLines(input)
	for _, line := range analyzeProtectedLines(lines) {
		if line.ContainsComment {
			return true
		}
	}
	return false
}

// SQLCommentLines returns complete source lines intersecting a lexical SQL
// comment. Keeping the whole line makes leading indentation and inline-comment
// placement part of the formatting safety contract.
func SQLCommentLines(input string) []string {
	lines := analyzeProtectedLines(splitSQLLines(input))
	out := make([]string, 0)
	for _, line := range lines {
		if line.ContainsComment {
			out = append(out, line.Raw)
		}
	}
	return out
}

func splitSQLLines(input string) []string {
	return strings.Split(strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\r", "\n"), "\n")
}

func ReindentTSQL(input string, indentSize int) string {
	if indentSize <= 0 {
		indentSize = 2
	}
	lines := strings.Split(normalizeSQL(input), "\n")
	protected := analyzeProtectedLines(lines)
	hasProtectedLines := false
	for _, line := range protected {
		if line.Preserve {
			hasProtectedLines = true
			break
		}
	}
	out := make([]string, 0, len(lines))
	stack := make([]contextTokenType, 0)
	nextLineExtraIndent := 0
	lastBlank := false
	clauseChain := clauseChainNone
	listContext := listContextNone
	selectCaseDepth := 0
	exprParenDepth := 0
	parenListContext := parenListNone
	pendingParenList := parenListNone
	queryParenDepth := 0

	for lineIndex, rawLine := range lines {
		if protected[lineIndex].Preserve {
			out = append(out, protected[lineIndex].Raw)
			lastBlank = false
			code := strings.TrimSpace(protected[lineIndex].Code)
			if code == "" {
				continue
			}
			tokens := scanWordTokens(code)
			inCaseExpression := caseDepth(stack) > 0
			inSelectCaseProjection := listContext == listContextSelect && selectCaseDepth > 0
			inProjectionCase := inCaseExpression || inSelectCaseProjection
			nextLineExtraIndent = 0
			stack = applyContextTokens(stack, tokens)
			clauseChain = updateClauseChain(tokens, clauseChain)
			listContext = updateListContext(code, tokens, listContext, inProjectionCase)
			if listContext == listContextSelect {
				for _, tok := range tokens {
					switch strings.ToUpper(tok.Text) {
					case "CASE":
						selectCaseDepth++
					case "END":
						if selectCaseDepth > 0 {
							selectCaseDepth--
						}
					}
				}
			} else if selectCaseDepth > 0 {
				selectCaseDepth = 0
			}
			exprParenDepth = updateParenDepth(exprParenDepth, code)
			parenListContext = updateParenListContext(code, tokens, parenListContext)
			if parenListContext == parenListNone {
				pendingParenList = detectPendingParenList(code, tokens)
			}
			queryParenDepth = updateQueryParenDepth(queryParenDepth, code, tokens, parenListContext)
			if startsWithSingleStatementControl(tokens) && !lineHasBlockBeginToken(tokens) {
				nextLineExtraIndent = 1
			}
			continue
		}
		expanded := ExpandTrailingControlBegin(strings.TrimSpace(rawLine))
		for _, line := range expanded {
			caseExpanded := splitCaseHeavyLine(line, caseDepth(stack) > 0)
			for _, subLine := range caseExpanded {
				for _, endAdjusted := range SplitEndTrailingSuffix(subLine) {
					line = endAdjusted
					if line == "" {
						if pendingParenList != parenListNone {
							// Keep CREATE/INSERT header directly followed by "(" list.
							nextLineExtraIndent = 0
							continue
						}
						nextLineExtraIndent = 0
						if len(out) > 0 && !lastBlank {
							out = append(out, "")
							lastBlank = true
						}
						continue
					}
					lastBlank = false

					tokens := scanWordTokens(line)
					trimmedLine := strings.TrimSpace(line)
					if parenListContext == parenListNone {
						inlineCtx := detectInlineParenList(trimmedLine, tokens)
						if inlineCtx != parenListNone {
							parenListContext = inlineCtx
						}
					}
					line = normalizeLineForFormatting(trimmedLine, tokens, listContext)
					tokens = scanWordTokens(line)
					trimmedLine = strings.TrimSpace(line)
					if pendingParenList != parenListNone {
						if strings.HasPrefix(trimmedLine, "(") {
							parenListContext = pendingParenList
						}
						pendingParenList = parenListNone
					}
					preDedent := 0
					if startsWithToken(tokens, "END") && topContextIs(stack, contextBlock) {
						preDedent = 1
					}
					head := ""
					if len(tokens) > 0 {
						head = strings.ToUpper(tokens[0].Text)
					}
					inCaseExpression := caseDepth(stack) > 0
					inSelectCaseProjection := listContext == listContextSelect && selectCaseDepth > 0
					inProjectionCase := inCaseExpression || inSelectCaseProjection

					currentExtra := nextLineExtraIndent
					if startsWithToken(tokens, "ELSE") || startsWithToken(tokens, "END") {
						currentExtra = 0
					}
					if startsWithToken(tokens, "BEGIN") {
						currentExtra = 0
					}
					if currentExtra < 0 {
						currentExtra = 0
					}
					if inProjectionCase {
						// CASE projection indentation is fully controlled by CASE rules.
						currentExtra = 0
					}
					baseIndent := blockDepth(stack) - preDedent
					if baseIndent < 0 {
						baseIndent = 0
					}
					caseIndent := caseDepth(stack)
					if startsWithToken(tokens, "END") && topContextIs(stack, contextCase) && caseIndent > 0 {
						caseIndent--
					}
					if isCaseKeyword(head) {
						caseIndent = 0
					}
					if listContext == listContextSelect && inProjectionCase {
						// CASE projection alignment in SELECT list is handled by tab-based
						// list continuation to avoid mixed space+tab indentation.
						caseIndent = 0
					}
					lineIndent := strings.Repeat(" ", (baseIndent+currentExtra+caseIndent)*indentSize)
					lineIndent += clauseContinuationPrefix(tokens, clauseChain)
					lineIndent += listContinuationPrefix(line, tokens, listContext, inProjectionCase)
					lineIndent += expressionContinuationPrefix(line, tokens, listContext, exprParenDepth, indentSize, inProjectionCase)
					lineIndent += parenListContinuationPrefix(trimmedLine, tokens, parenListContext)
					lineIndent += queryParenContinuationPrefix(trimmedLine, tokens, queryParenDepth, indentSize, listContext)
					caseIndentBase := strings.Repeat(" ", (baseIndent+currentExtra)*indentSize)
					if head == "CASE" && listContext == listContextSelect {
						// Start CASE projection at first SELECT-list indent level.
						lineIndent = caseIndentBase + "\t"
					} else if inProjectionCase && (head == "WHEN" || head == "THEN" || (head == "ELSE" && !isElseIfToken(tokens))) {
						// Keep WHEN/THEN/ELSE branches aligned one level deeper.
						lineIndent = caseIndentBase + "\t\t"
					} else if (topContextIs(stack, contextCase) || isSelectCaseEndProjectionLine(trimmedLine, tokens)) && head == "END" {
						// Keep CASE END aligned with CASE in projection context.
						lineIndent = caseIndentBase + "\t"
					}
					out = append(out, lineIndent+line)

					nextLineExtraIndent = 0
					stack = applyContextTokens(stack, tokens)
					clauseChain = updateClauseChain(tokens, clauseChain)
					listContext = updateListContext(line, tokens, listContext, inProjectionCase)
					if listContext == listContextSelect {
						for _, tok := range tokens {
							switch strings.ToUpper(tok.Text) {
							case "CASE":
								selectCaseDepth++
							case "END":
								if selectCaseDepth > 0 {
									selectCaseDepth--
								}
							}
						}
					} else if selectCaseDepth > 0 {
						selectCaseDepth = 0
					}
					exprParenDepth = updateParenDepth(exprParenDepth, line)
					parenListContext = updateParenListContext(trimmedLine, tokens, parenListContext)
					if parenListContext == parenListNone {
						pendingParenList = detectPendingParenList(trimmedLine, tokens)
					}
					queryParenDepth = updateQueryParenDepth(queryParenDepth, trimmedLine, tokens, parenListContext)
					if startsWithSingleStatementControl(tokens) &&
						!lineHasBlockBeginToken(tokens) &&
						!isSelectCaseProjectionLine(trimmedLine, tokens, inProjectionCase) {
						nextLineExtraIndent = 1
					}
				}
			}
		}
	}

	if !hasProtectedLines {
		out = compactQueryBlankLines(out)
	}
	return normalizeSQL(strings.Join(out, "\n"))
}

func startsWithToken(tokens []wordToken, text string) bool {
	if len(tokens) == 0 {
		return false
	}
	return strings.EqualFold(tokens[0].Text, text)
}

func isElseIfToken(tokens []wordToken) bool {
	if len(tokens) < 2 {
		return false
	}
	return strings.EqualFold(tokens[0].Text, "ELSE") && strings.EqualFold(tokens[1].Text, "IF")
}

func topContextIs(stack []contextTokenType, target contextTokenType) bool {
	if len(stack) == 0 {
		return false
	}
	return stack[len(stack)-1] == target
}

func blockDepth(stack []contextTokenType) int {
	n := 0
	for _, item := range stack {
		if item == contextBlock {
			n++
		}
	}
	return n
}

func caseDepth(stack []contextTokenType) int {
	n := 0
	for _, item := range stack {
		if item == contextCase {
			n++
		}
	}
	return n
}

func applyContextTokens(stack []contextTokenType, tokens []wordToken) []contextTokenType {
	for i := 0; i < len(tokens); i++ {
		word := strings.ToUpper(tokens[i].Text)
		switch word {
		case "CASE":
			stack = append(stack, contextCase)
		case "BEGIN":
			if isBeginNonBlock(tokens, i) {
				continue
			}
			stack = append(stack, contextBlock)
		case "END":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack
}

func lineHasBlockBeginToken(tokens []wordToken) bool {
	for i := 0; i < len(tokens); i++ {
		if !strings.EqualFold(tokens[i].Text, "BEGIN") {
			continue
		}
		if !isBeginNonBlock(tokens, i) {
			return true
		}
	}
	return false
}

func isBeginNonBlock(tokens []wordToken, idx int) bool {
	if idx+1 >= len(tokens) {
		return false
	}
	next := strings.ToUpper(tokens[idx+1].Text)
	switch next {
	case "TRAN", "TRANSACTION", "DIALOG", "CONVERSATION", "DISTRIBUTED":
		return true
	default:
		return false
	}
}

func startsWithSingleStatementControl(tokens []wordToken) bool {
	if len(tokens) == 0 {
		return false
	}
	t0 := strings.ToUpper(tokens[0].Text)
	switch t0 {
	case "IF", "WHILE":
		return true
	case "ELSE":
		if len(tokens) == 1 {
			return true
		}
		return strings.EqualFold(tokens[1].Text, "IF")
	default:
		return false
	}
}

func ExpandTrailingControlBegin(line string) []string {
	line = strings.TrimSpace(line)
	if line == "" {
		return []string{""}
	}
	lineNoSemi := strings.TrimSpace(strings.TrimSuffix(line, ";"))
	tokens := scanWordTokens(lineNoSemi)
	if len(tokens) == 0 {
		return []string{line}
	}
	if !startsWithSingleStatementControl(tokens) {
		return []string{line}
	}
	last := tokens[len(tokens)-1]
	if !strings.EqualFold(last.Text, "BEGIN") {
		return []string{line}
	}
	if last.End != len(lineNoSemi) {
		return []string{line}
	}
	control := strings.TrimSpace(lineNoSemi[:last.Start])
	if control == "" || strings.EqualFold(control, "BEGIN") {
		return []string{line}
	}
	return []string{control, "BEGIN"}
}

func scanWordTokens(line string) []wordToken {
	tokens := make([]wordToken, 0)
	inString := false
	inBracket := false
	i := 0
	for i < len(line) {
		ch := line[i]
		next := byte(0)
		if i+1 < len(line) {
			next = line[i+1]
		}

		if inString {
			if ch == '\'' {
				if next == '\'' {
					i += 2
					continue
				}
				inString = false
			}
			i++
			continue
		}
		if inBracket {
			if ch == ']' {
				if next == ']' {
					i += 2
					continue
				}
				inBracket = false
			}
			i++
			continue
		}
		if ch == '\'' {
			inString = true
			i++
			continue
		}
		if ch == '[' {
			inBracket = true
			i++
			continue
		}
		if isWordStart(ch) {
			start := i
			i++
			for i < len(line) && isWordPart(line[i]) {
				i++
			}
			tokens = append(tokens, wordToken{
				Text:  line[start:i],
				Start: start,
				End:   i,
			})
			continue
		}
		i++
	}
	return tokens
}

func splitCaseHeavyLine(line string, activeCase bool) []string {
	line = strings.TrimSpace(line)
	if line == "" {
		return []string{""}
	}
	tokens := scanWordTokens(line)
	if len(tokens) < 2 {
		return []string{line}
	}
	caseCount := 0
	caseKeywordTotal := 0
	for _, tok := range tokens {
		switch strings.ToUpper(tok.Text) {
		case "CASE":
			caseCount++
			caseKeywordTotal++
		case "WHEN", "ELSE", "END":
			caseKeywordTotal++
		}
	}
	if caseKeywordTotal < 2 {
		return []string{line}
	}
	if caseCount == 0 && !activeCase {
		return []string{line}
	}
	if caseCount == 0 && caseKeywordTotal < 2 {
		return []string{line}
	}

	breakStarts := []int{0}
	for i := 1; i < len(tokens); i++ {
		word := strings.ToUpper(tokens[i].Text)
		shouldBreak := false
		switch word {
		case "CASE", "WHEN", "ELSE", "END":
			shouldBreak = true
		}
		if !shouldBreak {
			continue
		}
		if word == "ELSE" && i+1 < len(tokens) && strings.EqualFold(tokens[i+1].Text, "IF") {
			continue
		}
		start := tokens[i].Start
		if start <= 0 || start >= len(line) {
			continue
		}
		last := breakStarts[len(breakStarts)-1]
		if start == last {
			continue
		}
		breakStarts = append(breakStarts, start)
	}
	if len(breakStarts) == 1 {
		return []string{line}
	}

	parts := make([]string, 0, len(breakStarts))
	for i := 0; i < len(breakStarts); i++ {
		start := breakStarts[i]
		end := len(line)
		if i+1 < len(breakStarts) {
			end = breakStarts[i+1]
		}
		part := strings.TrimSpace(line[start:end])
		if part == "" {
			continue
		}
		parts = append(parts, part)
	}
	if len(parts) == 0 {
		return []string{line}
	}
	return parts
}

func SplitEndTrailingSuffix(line string) []string {
	line = strings.TrimSpace(line)
	if line == "" {
		return []string{""}
	}
	tokens := scanWordTokens(line)
	if len(tokens) == 0 || !strings.EqualFold(tokens[0].Text, "END") {
		return []string{line}
	}
	if tokens[0].End >= len(line) {
		return []string{line}
	}
	rest := strings.TrimSpace(line[tokens[0].End:])
	if rest == "" {
		return []string{"END"}
	}
	restUpper := strings.ToUpper(rest)
	if strings.HasPrefix(restUpper, "AS ") {
		return []string{line}
	}
	if strings.HasPrefix(rest, ")") {
		return []string{"END", rest}
	}
	return []string{line}
}

func isWordStart(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_'
}

func isWordPart(ch byte) bool {
	return isWordStart(ch) || (ch >= '0' && ch <= '9')
}

type clauseChainType int

const (
	clauseChainNone clauseChainType = iota
	clauseChainOn
	clauseChainWhere
)

func clauseContinuationPrefix(tokens []wordToken, chain clauseChainType) string {
	if len(tokens) == 0 {
		return ""
	}
	switch strings.ToUpper(tokens[0].Text) {
	case "ON":
		return "\t"
	case "AND":
		if chain == clauseChainOn {
			return "\t\t"
		}
		if chain == clauseChainWhere {
			return "\t"
		}
	}
	return ""
}

func updateClauseChain(tokens []wordToken, current clauseChainType) clauseChainType {
	if len(tokens) == 0 {
		return current
	}
	head := strings.ToUpper(tokens[0].Text)
	switch head {
	case "ON":
		return clauseChainOn
	case "WHERE":
		return clauseChainWhere
	case "AND":
		return current
	case "JOIN", "LEFT", "RIGHT", "INNER", "FULL", "CROSS", "FROM", "ORDER", "GROUP", "HAVING":
		return clauseChainNone
	default:
		return clauseChainNone
	}
}

type listContextType int

const (
	listContextNone listContextType = iota
	listContextSelect
	listContextDeclare
	listContextUpdateSet
	listContextProcedureParams
)

func listContinuationPrefix(line string, tokens []wordToken, ctx listContextType, inCaseExpression bool) string {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return ""
	}
	switch ctx {
	case listContextSelect:
		if strings.HasPrefix(trimmed, ",") {
			return "\t"
		}
		if len(tokens) == 0 {
			return "\t"
		}
		head := strings.ToUpper(tokens[0].Text)
		if inCaseExpression {
			switch head {
			case "CASE", "END":
				return "\t"
			case "WHEN", "THEN", "ELSE":
				return "\t\t"
			default:
				if isSelectListTerminator(head) {
					return ""
				}
				// Keep expression payload lines inside CASE one level deeper
				// than WHEN/THEN/ELSE for better readability.
				return "\t\t\t"
			}
		}
		if isStatementBoundary(head) && !isSelectCaseProjectionLine(trimmed, tokens, inCaseExpression) {
			return ""
		}
		if isSelectListTerminator(head) {
			return ""
		}
		return "\t"
	case listContextDeclare:
		if strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @") {
			return "\t"
		}
	case listContextProcedureParams:
		if strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @") {
			return "\t"
		}
	case listContextUpdateSet:
		if len(tokens) == 0 {
			return ""
		}
		head := strings.ToUpper(tokens[0].Text)
		if head == "SET" || isUpdateSetTerminator(head) || isStatementBoundary(head) {
			return ""
		}
		return "\t"
	}
	return ""
}

func updateListContext(line string, tokens []wordToken, current listContextType, inCaseExpression bool) listContextType {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return current
	}
	if len(tokens) > 0 {
		switch strings.ToUpper(tokens[0].Text) {
		case "SELECT":
			if strings.Contains(trimmed, ";") {
				return listContextNone
			}
			return listContextSelect
		case "DECLARE":
			if strings.Contains(trimmed, ";") {
				return listContextNone
			}
			return listContextDeclare
		case "SET":
			if strings.Contains(trimmed, ";") {
				return listContextNone
			}
			return listContextUpdateSet
		}
		if isProcedureHeaderStart(tokens) {
			if hasStandaloneAsToken(tokens) {
				return listContextNone
			}
			return listContextProcedureParams
		}
	}
	switch current {
	case listContextSelect:
		if strings.Contains(trimmed, ";") {
			return listContextNone
		}
		if len(tokens) == 0 {
			return current
		}
		head := strings.ToUpper(tokens[0].Text)
		if isSelectListTerminator(head) {
			return listContextNone
		}
		if isStatementBoundary(head) && !isSelectCaseProjectionLine(trimmed, tokens, inCaseExpression) {
			return listContextNone
		}
		return current
	case listContextDeclare:
		if strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @") {
			return current
		}
		return listContextNone
	case listContextUpdateSet:
		if strings.Contains(trimmed, ";") {
			return listContextNone
		}
		if len(tokens) == 0 {
			return current
		}
		head := strings.ToUpper(tokens[0].Text)
		if isUpdateSetTerminator(head) || isStatementBoundary(head) {
			return listContextNone
		}
		return current
	case listContextProcedureParams:
		if len(tokens) == 0 {
			return current
		}
		head := strings.ToUpper(tokens[0].Text)
		if head == "AS" {
			return listContextNone
		}
		if strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @") {
			return current
		}
		if head == "WITH" {
			return current
		}
		return listContextNone
	default:
		return listContextNone
	}
}

func expressionContinuationPrefix(line string, tokens []wordToken, listCtx listContextType, depth int, indentSize int, inCaseExpression bool) string {
	if listCtx != listContextSelect {
		return ""
	}
	if inCaseExpression {
		return ""
	}
	if depth <= 0 {
		return ""
	}
	if len(tokens) > 0 && isSelectListTerminator(strings.ToUpper(tokens[0].Text)) {
		return ""
	}
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return ""
	}
	if !shouldApplyExpressionContinuation(trimmed, tokens, inCaseExpression) {
		return ""
	}
	effectiveDepth := depth - leadingCloseParenCount(trimmed)
	if effectiveDepth < 0 {
		effectiveDepth = 0
	}
	if effectiveDepth == 0 {
		return ""
	}
	if effectiveDepth > 2 {
		effectiveDepth = 2
	}
	return strings.Repeat(" ", effectiveDepth*indentSize)
}

func shouldApplyExpressionContinuation(trimmed string, tokens []wordToken, inCaseExpression bool) bool {
	if strings.HasPrefix(trimmed, "(") || strings.HasPrefix(trimmed, ")") {
		return true
	}
	if len(tokens) == 0 {
		return false
	}
	head := strings.ToUpper(tokens[0].Text)
	if head == "CASE" {
		return false
	}
	if inCaseExpression {
		switch head {
		case "WHEN", "THEN", "ELSE", "END":
			// CASE indentation is handled by caseDepth stack.
			return false
		}
	}
	switch head {
	case "WHEN", "THEN", "ELSE", "END":
		return true
	default:
		return false
	}
}

func updateParenDepth(current int, line string) int {
	delta := parenDeltaIgnoringStrings(line)
	next := current + delta
	if next < 0 {
		next = 0
	}
	return next
}

func leadingCloseParenCount(trimmed string) int {
	count := 0
	for i := 0; i < len(trimmed); i++ {
		if trimmed[i] == ')' {
			count++
			continue
		}
		if trimmed[i] == ' ' || trimmed[i] == '\t' {
			continue
		}
		break
	}
	return count
}

func parenDeltaIgnoringStrings(line string) int {
	inString := false
	inBracket := false
	delta := 0
	for i := 0; i < len(line); i++ {
		ch := line[i]
		next := byte(0)
		if i+1 < len(line) {
			next = line[i+1]
		}

		if inString {
			if ch == '\'' {
				if next == '\'' {
					i++
					continue
				}
				inString = false
			}
			continue
		}
		if inBracket {
			if ch == ']' {
				if next == ']' {
					i++
					continue
				}
				inBracket = false
			}
			continue
		}
		if ch == '\'' {
			inString = true
			continue
		}
		if ch == '[' {
			inBracket = true
			continue
		}
		if ch == '(' {
			delta++
			continue
		}
		if ch == ')' {
			delta--
		}
	}
	return delta
}

func isSelectListTerminator(head string) bool {
	switch head {
	case "FROM", "INTO", "WHERE", "GROUP", "HAVING", "ORDER", "UNION", "EXCEPT", "INTERSECT", "OPTION":
		return true
	default:
		return false
	}
}

func isStatementBoundary(head string) bool {
	switch head {
	case "IF", "ELSE", "BEGIN", "END", "DECLARE", "SET", "UPDATE", "INSERT", "DELETE", "MERGE", "CREATE", "ALTER", "DROP", "TRUNCATE", "RETURN", "EXEC", "EXECUTE", "WHILE", "PRINT", "RAISERROR", "THROW":
		return true
	default:
		return false
	}
}

func isUpdateSetTerminator(head string) bool {
	switch head {
	case "FROM", "WHERE", "OUTPUT", "OPTION", "WHEN", "VALUES", "USING", "MERGE", "INSERT", "DELETE", "UPDATE":
		return true
	default:
		return false
	}
}

func compactQueryBlankLines(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	out := make([]string, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) != "" {
			out = append(out, line)
			continue
		}
		prev := previousNonBlank(lines, i-1)
		next := nextNonBlank(lines, i+1)
		if prev == "" || next == "" {
			out = append(out, line)
			continue
		}
		if shouldCompactBlankBetween(prev, next) {
			continue
		}
		out = append(out, line)
	}
	return out
}

func previousNonBlank(lines []string, from int) string {
	for i := from; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			return strings.TrimSpace(lines[i])
		}
	}
	return ""
}

func nextNonBlank(lines []string, from int) string {
	for i := from; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) != "" {
			return strings.TrimSpace(lines[i])
		}
	}
	return ""
}

func shouldCompactBlankBetween(prevLine, nextLine string) bool {
	prevLine = strings.TrimSpace(prevLine)
	nextLine = strings.TrimSpace(nextLine)
	prevHead := firstTokenUpper(prevLine)
	nextHead := firstTokenUpper(nextLine)
	// Keep SELECT projection list compact (comma-separated items),
	// including when next item starts with literal/quoted/bracketed token.
	if strings.HasSuffix(prevLine, ",") && nextLine != "" {
		return true
	}
	if nextHead == "" {
		return false
	}
	// Keep visual separation before new top-level statements.
	if isStatementBoundary(nextHead) {
		return false
	}
	if isProcedureHeaderLine(prevLine) && isProcedureParamLine(nextLine) {
		return true
	}
	if isProcedureParamLine(prevLine) && isProcedureParamLine(nextLine) {
		return true
	}
	if isDeclareListLine(prevLine) && isDeclareListLine(nextLine) {
		return true
	}
	if strings.HasPrefix(nextLine, ")") && isCaseKeyword(prevHead) {
		return true
	}
	if isCaseKeyword(nextHead) {
		if prevLine == "(" || prevLine == ")" {
			return true
		}
		if strings.HasSuffix(prevLine, "+") || strings.HasSuffix(prevLine, ",") {
			return true
		}
		if isCaseKeyword(prevHead) {
			return true
		}
	}
	// Remove blank lines that split a single query chain.
	if isQueryContinuationHead(nextHead) {
		return true
	}
	// Remove blank lines between SELECT-list/SET-list and their FROM/WHERE continuation.
	if prevHead == "SELECT" || prevHead == "SET" || strings.HasPrefix(prevLine, ",") {
		if isQueryContinuationHead(nextHead) {
			return true
		}
	}
	return false
}

func firstTokenUpper(line string) string {
	tokens := scanWordTokens(strings.TrimSpace(line))
	if len(tokens) == 0 {
		return ""
	}
	return strings.ToUpper(tokens[0].Text)
}

func isQueryContinuationHead(head string) bool {
	switch head {
	case "FROM", "JOIN", "LEFT", "RIGHT", "INNER", "FULL", "CROSS", "WHERE", "GROUP", "HAVING", "ORDER", "ON", "AND", "OR", "UNION", "EXCEPT", "INTERSECT", "INTO":
		return true
	default:
		return false
	}
}

func isCaseKeyword(head string) bool {
	switch head {
	case "CASE", "WHEN", "THEN", "ELSE", "END":
		return true
	default:
		return false
	}
}

func isSelectCaseEndProjectionLine(trimmed string, tokens []wordToken) bool {
	if len(tokens) == 0 {
		return false
	}
	if !strings.EqualFold(tokens[0].Text, "END") {
		return false
	}
	if tokens[0].End > len(trimmed) {
		return false
	}
	rest := strings.TrimSpace(trimmed[tokens[0].End:])
	if rest == "" || rest == ";" {
		return false
	}
	upper := strings.ToUpper(rest)
	if strings.HasPrefix(upper, "AS ") {
		return true
	}
	if strings.Contains(rest, ",") {
		return true
	}
	return false
}

func isSelectCaseProjectionLine(trimmed string, tokens []wordToken, inCaseExpression bool) bool {
	if len(tokens) == 0 {
		return false
	}
	head := strings.ToUpper(tokens[0].Text)
	switch head {
	case "CASE":
		return true
	case "WHEN":
		return inCaseExpression
	case "ELSE":
		if len(tokens) > 1 && strings.EqualFold(tokens[1].Text, "IF") {
			return false
		}
		return inCaseExpression
	case "END":
		if isSelectCaseEndProjectionLine(trimmed, tokens) {
			return true
		}
		return inCaseExpression
	default:
		return false
	}
}

type parenListType int

const (
	parenListNone parenListType = iota
	parenListCreateTable
	parenListInsert
)

func parenListContinuationPrefix(trimmed string, tokens []wordToken, ctx parenListType) string {
	if ctx == parenListNone {
		return ""
	}
	if len(tokens) > 0 && isStatementBoundary(strings.ToUpper(tokens[0].Text)) {
		return ""
	}
	if trimmed == "" || trimmed == "(" || strings.HasPrefix(trimmed, ")") {
		return ""
	}
	return "\t"
}

func updateParenListContext(trimmed string, tokens []wordToken, current parenListType) parenListType {
	if current == parenListNone {
		return parenListNone
	}
	if trimmed == ")" || strings.HasPrefix(trimmed, ");") {
		return parenListNone
	}
	// Keep inside list until closing parenthesis.
	return current
}

func detectPendingParenList(trimmed string, tokens []wordToken) parenListType {
	if len(tokens) == 0 {
		return parenListNone
	}
	upper := strings.ToUpper(strings.Join(tokenTexts(tokens), " "))
	if strings.HasPrefix(upper, "CREATE TABLE") && !strings.Contains(trimmed, "(") {
		return parenListCreateTable
	}
	if strings.HasPrefix(upper, "INSERT INTO") && !strings.Contains(trimmed, "(") {
		return parenListInsert
	}
	if strings.HasPrefix(upper, "INSERT") && !strings.Contains(trimmed, "(") {
		return parenListInsert
	}
	if strings.HasPrefix(upper, "VALUES") && !strings.Contains(trimmed, "(") {
		return parenListInsert
	}
	if isDeclareTableHeader(trimmed, tokens) && !strings.Contains(trimmed, "(") {
		return parenListCreateTable
	}
	return parenListNone
}

func detectInlineParenList(trimmed string, tokens []wordToken) parenListType {
	if len(tokens) == 0 || !strings.Contains(trimmed, "(") {
		return parenListNone
	}
	openIdx := strings.Index(trimmed, "(")
	closeIdx := strings.LastIndex(trimmed, ")")
	if openIdx >= 0 && closeIdx > openIdx {
		// Inline parenthesis already closed in the same line, do not keep list context.
		return parenListNone
	}
	upper := strings.ToUpper(strings.Join(tokenTexts(tokens), " "))
	if strings.HasPrefix(upper, "CREATE TABLE") {
		return parenListCreateTable
	}
	if strings.HasPrefix(upper, "INSERT INTO") {
		return parenListInsert
	}
	if strings.HasPrefix(upper, "INSERT") {
		return parenListInsert
	}
	if strings.HasPrefix(upper, "VALUES") {
		return parenListInsert
	}
	if isDeclareTableHeader(trimmed, tokens) {
		return parenListCreateTable
	}
	return parenListNone
}

func isDeclareTableHeader(trimmed string, tokens []wordToken) bool {
	if len(tokens) == 0 {
		return false
	}
	if !strings.EqualFold(tokens[0].Text, "DECLARE") {
		return false
	}
	upper := strings.ToUpper(trimmed)
	if strings.Contains(upper, " TABLE") {
		return true
	}
	return false
}

func queryParenContinuationPrefix(trimmed string, tokens []wordToken, depth int, indentSize int, listCtx listContextType) string {
	if depth <= 0 {
		return ""
	}
	if strings.HasPrefix(trimmed, ")") {
		return ""
	}
	if listCtx == listContextSelect || listCtx == listContextDeclare || listCtx == listContextUpdateSet {
		return ""
	}
	if len(tokens) == 0 {
		return ""
	}
	head := strings.ToUpper(tokens[0].Text)
	if head == "ON" || head == "AND" || head == "OR" {
		return ""
	}
	if !isQueryContinuationHead(head) && head != "SELECT" && head != "WITH" {
		return ""
	}
	return "\t"
}

func updateQueryParenDepth(current int, trimmed string, tokens []wordToken, parenListCtx parenListType) int {
	depth := current
	if strings.HasPrefix(trimmed, ")") && depth > 0 {
		depth--
	}
	if parenListCtx != parenListNone {
		return depth
	}
	if isQueryParenOpener(trimmed, tokens) {
		depth++
	}
	return depth
}

func isQueryParenOpener(trimmed string, tokens []wordToken) bool {
	if !strings.HasSuffix(trimmed, "(") {
		return false
	}
	if strings.Contains(strings.ToUpper(trimmed), " AS (") {
		return true
	}
	if len(tokens) == 0 {
		return false
	}
	switch strings.ToUpper(tokens[0].Text) {
	case "SELECT", "FROM", "JOIN", "LEFT", "RIGHT", "INNER", "FULL", "CROSS", "OUTER", "USING", "VALUES", "WITH":
		return true
	default:
		return false
	}
}

func tokenTexts(tokens []wordToken) []string {
	out := make([]string, len(tokens))
	for i, tok := range tokens {
		out[i] = tok.Text
	}
	return out
}

func normalizeLineForFormatting(trimmed string, tokens []wordToken, ctx listContextType) string {
	if trimmed == "" || len(tokens) == 0 {
		return trimmed
	}
	head := strings.ToUpper(tokens[0].Text)
	switch head {
	case "DECLARE":
		if normalized, ok := normalizeDeclareHeaderLine(trimmed, tokens[0].End); ok {
			return normalized
		}
	}
	switch ctx {
	case listContextDeclare, listContextProcedureParams:
		if normalized, ok := normalizeParamLine(trimmed); ok {
			return normalized
		}
	}
	return trimmed
}

func normalizeDeclareHeaderLine(trimmed string, declareTokenEnd int) (string, bool) {
	if declareTokenEnd <= 0 || declareTokenEnd > len(trimmed) {
		return "", false
	}
	declarePrefix := trimmed[:declareTokenEnd]
	rest := strings.TrimSpace(trimmed[declareTokenEnd:])
	if rest == "" {
		return "", false
	}
	normalizedRest, ok := normalizeParamTail(rest)
	if !ok {
		return "", false
	}
	return declarePrefix + " " + normalizedRest, true
}

func normalizeParamLine(trimmed string) (string, bool) {
	normalized, ok := normalizeParamTail(trimmed)
	if !ok {
		return "", false
	}
	return normalized, true
}

func normalizeParamTail(input string) (string, bool) {
	s := strings.TrimSpace(input)
	if s == "" {
		return "", false
	}

	hasLeadingComma := false
	if strings.HasPrefix(s, ",") {
		hasLeadingComma = true
		s = strings.TrimSpace(s[1:])
	}
	if !strings.HasPrefix(s, "@") {
		return "", false
	}

	name, rest := splitFirstWhitespace(s)
	if name == "" {
		return "", false
	}

	if rest == "" {
		if hasLeadingComma {
			return ", " + name, true
		}
		return name, true
	}

	rest = strings.TrimSpace(compactHorizontalWhitespace(rest))
	if hasLeadingComma {
		return ", " + name + " " + rest, true
	}
	return name + " " + rest, true
}

func splitFirstWhitespace(input string) (string, string) {
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' || input[i] == '\t' {
			return input[:i], strings.TrimSpace(input[i:])
		}
	}
	return input, ""
}

func compactHorizontalWhitespace(input string) string {
	var b strings.Builder
	b.Grow(len(input))

	inString := false
	inBracket := false
	lastWasSpace := false
	for i := 0; i < len(input); i++ {
		ch := input[i]
		next := byte(0)
		if i+1 < len(input) {
			next = input[i+1]
		}

		if inString {
			b.WriteByte(ch)
			if ch == '\'' {
				if next == '\'' {
					b.WriteByte(next)
					i++
					continue
				}
				inString = false
			}
			lastWasSpace = false
			continue
		}
		if inBracket {
			b.WriteByte(ch)
			if ch == ']' {
				if next == ']' {
					b.WriteByte(next)
					i++
					continue
				}
				inBracket = false
			}
			lastWasSpace = false
			continue
		}

		if ch == '\'' {
			inString = true
			b.WriteByte(ch)
			lastWasSpace = false
			continue
		}
		if ch == '[' {
			inBracket = true
			b.WriteByte(ch)
			lastWasSpace = false
			continue
		}
		if ch == ' ' || ch == '\t' {
			if !lastWasSpace {
				b.WriteByte(' ')
				lastWasSpace = true
			}
			continue
		}

		b.WriteByte(ch)
		lastWasSpace = false
	}
	return strings.TrimSpace(b.String())
}

func isProcedureHeaderStart(tokens []wordToken) bool {
	if len(tokens) < 2 {
		return false
	}
	first := strings.ToUpper(tokens[0].Text)
	switch first {
	case "CREATE":
		i := 1
		if len(tokens) > 3 && strings.EqualFold(tokens[1].Text, "OR") && strings.EqualFold(tokens[2].Text, "ALTER") {
			i = 3
		}
		if i < len(tokens) {
			head := strings.ToUpper(tokens[i].Text)
			return head == "PROCEDURE" || head == "PROC"
		}
	case "ALTER":
		if len(tokens) > 1 {
			head := strings.ToUpper(tokens[1].Text)
			return head == "PROCEDURE" || head == "PROC"
		}
	}
	return false
}

func hasStandaloneAsToken(tokens []wordToken) bool {
	for _, tok := range tokens {
		if strings.EqualFold(tok.Text, "AS") {
			return true
		}
	}
	return false
}

func isProcedureHeaderLine(line string) bool {
	tokens := scanWordTokens(strings.TrimSpace(line))
	return isProcedureHeaderStart(tokens)
}

func isProcedureParamLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @")
}

func isDeclareListLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(strings.ToUpper(trimmed), "DECLARE @") {
		return true
	}
	return strings.HasPrefix(trimmed, "@") || strings.HasPrefix(trimmed, ",@") || strings.HasPrefix(trimmed, ", @")
}
