package numrecipes

func GCD(number1, number2 uint) uint {
	for number2 != 0 {
		number1, number2 = number2, number1%number2
	}

	return number1
}
