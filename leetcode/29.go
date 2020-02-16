package main

import "fmt"

func main() {
	dividend, divisor := -2147483648, -1
	fmt.Println(divide(dividend, divisor))
	fmt.Println(1 << 31)
	fmt.Println(-1 * (1 << 31))
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	negative := (dividend ^ divisor) < 0
	dividend = abs(dividend)
	divisor = abs(divisor)
	var res int
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			res += 1 << i
			dividend -= divisor << i
		}
	}

	if negative {
		res = res * (-1)
	}
	if res < (-1*(1<<31)) || res > ((1<<31)-1) {
		return (1 << 31) - 1
	}
	return res
}

func abs(a int) int {
	if (a >> 31) == 0 {
		return a
	}
	return ^a + 1
}
