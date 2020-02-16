package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	last, res := 0.0, 1.0
	fx := float64(x)
	for res != last {
		last = res
		res = (res + fx/res) / 2
	}
	return int(fx)
}
