package piscine

func ToUpper(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// Convert lowercase letters to uppercase
			result += string(char - 32)
		} else {
			// Append non-letter characters as they are
			result += string(char)
		}
	}

	return result
}
