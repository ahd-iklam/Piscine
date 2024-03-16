package piscine

func Unmatch(a []int) int {
	// Create a map to store the frequency of each number
	frequency := make(map[int]int)

	// Iterate through the slice and count the frequency of each number
	for _, num := range a {
		frequency[num]++
	}

	// Iterate through the slice again to find the unpaired element
	for _, num := range a {
		if frequency[num] == 1 {
			return num
		}
	}

	// If no unpaired element found, return -1
	return -1
}
