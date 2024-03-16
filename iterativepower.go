package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		result := 1
		for i := 1; i <= power; i++ {
			result = nb * result
		}
		return result
	}
}
