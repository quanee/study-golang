package main

import "fmt"

func main() {
	arr := []int{3, 1, 7, 11}
	fmt.Println(checkIfExist(arr))
	fmt.Println(1 << 2)
	fmt.Println(2 << 2)
	fmt.Println(3 << 2)
}

func checkIfExist(arr []int) bool {
	for i, k := range arr {
		for j := i + 1; j < len(arr); j++ {
			if k<<1 == arr[j] || arr[j]<<1 == k {
				return true
			}
		}
	}
	return false
}
