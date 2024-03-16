package piscine

func CollatzCountdown(start int) int {
	counter := 0
	return Compteur(counter, start)
}

func Compteur(counter int, start int) int {
	if counter >= 5000 || start <= 0 {
		return -1
	}

	if start == 1 {
		return counter
	}

	if start%2 == 0 {
		start = start / 2
		counter++
		return Compteur(counter, start)
	} else {
		start = (start * 3) + 1
		counter++
		return Compteur(counter, start)
	}
}
