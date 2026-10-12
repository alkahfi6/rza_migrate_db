package sqlparse

// SQLToken is the read-only lexical view used for reporting. All consumers use
// the same scanner as object access/dependency discovery; literals remain opaque.
type SQLToken struct {
	Text, Upper        string
	Start, End, Depth  int
	Identifier, Quoted bool
}

func SQLTokens(sql string) []SQLToken {
	source := sourceTokens(sql)
	out := make([]SQLToken, len(source))
	for i, t := range source {
		out[i] = SQLToken{Text: t.text, Upper: t.upper, Start: t.start, End: t.end, Depth: t.depth, Identifier: t.identifier, Quoted: t.quoted}
	}
	return out
}
