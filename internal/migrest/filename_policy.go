package migrest

func migrationAllowedCategories() map[string]struct{} {
	return map[string]struct{}{
		"types":      {},
		"tables":     {},
		"indexes":    {},
		"views":      {},
		"functions":  {},
		"procedures": {},
		"patch":      {},
	}
}

func migrationAllowedObjectTokenRunes() map[rune]struct{} {
	out := make(map[rune]struct{})
	for ch := 'a'; ch <= 'z'; ch++ {
		out[ch] = struct{}{}
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		out[ch] = struct{}{}
	}
	for ch := '0'; ch <= '9'; ch++ {
		out[ch] = struct{}{}
	}
	out['_'] = struct{}{}
	out['#'] = struct{}{}
	out['.'] = struct{}{}
	return out
}
