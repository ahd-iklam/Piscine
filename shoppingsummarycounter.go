package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	// Initialize variables to keep track of word boundaries
	start := 0
	end := 0

	// Iterate over each character in the string
	for i := 0; i < len(str); i++ {
		// Check if the current character is a space or it's the end of the string
		if str[i] == ' ' || i == len(str)-1 {
			// Update the end index to the current position if it's not a space
			if i == len(str)-1 && str[i] != ' ' {
				end = i + 1
			} else {
				end = i
			}

			// Extract the word from the string
			word := str[start:end]

			// Increment the count for the word in the summary map
			summary[word]++

			// Move the start index to the next character after the space
			start = i + 1
		}
	}

	return summary
}
