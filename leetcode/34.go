package main

import "fmt"

func main() {
	nums := []int{2, 8, 8, 8, 9}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1
	if len(nums) < 2 {
		if len(nums) == 1 && nums[0] == target {
			return []int{0, 0}
		}
		return res
	}
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if target != nums[left] {
		return res
	}
	res[0] = left
	right = len(nums)
	for left < right {
		mid = left + (right-left)>>1
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	res[1] = left - 1

	return res
}
