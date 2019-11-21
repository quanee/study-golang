package main

import "fmt"

func test(n int) (x int) {
	defer func() {
		x += n
	}()
	/*defer func(x *int) {
		*x += n
	}(&x)
	defer func(x int) {
		x += n
	}(x)*/

	return n
}

func main() {
	fmt.Println(test(2))
}
