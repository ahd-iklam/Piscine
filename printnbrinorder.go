package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}

	sliceDigits := []int{}
	for n > 0 {
		md := n % 10
		sliceDigits = append(sliceDigits, md)
		n /= 10
	}
	for i := 0; i < len(sliceDigits); i++ {
		for j := i + 1; j < len(sliceDigits); j++ {
			if sliceDigits[i] > sliceDigits[j] {
				sliceDigits[i], sliceDigits[j] = sliceDigits[j], sliceDigits[i]
			}
		}
	}
	for _, digit := range sliceDigits {
		z01.PrintRune(rune(digit) + '0')
	}
}
