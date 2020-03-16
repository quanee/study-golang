package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 5, 6, 7, 8, 9, 12, 23, 34, 45, 56, 67, 78, 89, 90}
	fmt.Println(len(nums))
	i, m := (len(nums)+1)>>1, len(nums)>>1
	k := 90
	count := 0
	for m != 0 {
		if k < nums[i] {
			i = i - (m+1)>>1
			m = m >> 1
		} else if k > nums[i] {
			i = i + (m+1)>>1
			m = (m + 1) >> 1
		} else if k == nums[i] {
			break
		}
		count++
	}
	fmt.Println(nums[i], i)
	fmt.Println(count)
}
