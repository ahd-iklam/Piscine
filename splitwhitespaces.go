package piscine

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := 0
	for i, r := range s {
		if isWhiteSpace(r) {
			if i > start {
				words = append(words, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, s[start:])
	}
	return words
}
