package piscine

func StrRev(s string) string {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return string(runes)
}
