package main

import "fmt"

func main() {
	var a []int = nil
	a, a[0] = []int{1, 2}, 9
	fmt.Println(a)
}
