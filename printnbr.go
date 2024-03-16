package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}

	var digits [10]int
	idx := 0

	for n > 0 {
		digits[idx] = n % 10
		n /= 10
		idx++
	}

	for i := idx - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
