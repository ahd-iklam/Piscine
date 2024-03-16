package piscine

import (
	"strings"
)

func Capitalize(s string) string {
	words := strings.Fields(s) // Split the string into words
	var capitalizedWords []string

	for _, word := range words {
		if len(word) > 0 {
			// Capitalize the first letter and lowercase the rest of the word
			capitalizedWord := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			capitalizedWords = append(capitalizedWords, capitalizedWord)
		}
	}

	return strings.Join(capitalizedWords, " ")
}
