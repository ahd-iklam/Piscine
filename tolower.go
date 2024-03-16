package piscine

func ToLower(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			// Append non-letter characters as they are
			result += string(char)
		}
	}

	return result
}
