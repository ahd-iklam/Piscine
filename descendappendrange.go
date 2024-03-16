package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	num := []int{}

	for i := max; i > min; i-- {
		num = append(num, i)
	}

	return num
}
