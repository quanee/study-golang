package numrecipes

func QPow(number, x int) int {
	if x == 0 {
		return 1
	}
	if x%2 == 0 {
		return QPow(number*number, x/2)
	} else {
		return QPow(number, x-1) * number
	}
}
