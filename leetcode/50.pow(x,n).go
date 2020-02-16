package main

import "fmt"

func main() {
	fmt.Println(myPow(244, 3243))
}

func Pow(x float64, n int) float64 {
	return myPow(x, n)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPow(x*x, n>>1)
	} else {
		return myPow(x, n-1) * x
	}
}
