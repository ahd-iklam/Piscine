package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	for _, char := range s {
		if char == '-' && result == 0 {
			sign = -1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result * sign
}
