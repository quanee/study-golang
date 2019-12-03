package main

import (
	"fmt"
)

var debug = 2

func main() {
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = i * i
	}
	fmt.Println(nums)
	fmt.Println(debug)
}
