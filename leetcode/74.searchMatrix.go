package main

import "fmt"

func main() {
	mat := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	tag := 11
	fmt.Println(searchMatrix(mat, tag))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	left, right, mid := 0, len(matrix), 0
	for left < right {
		mid = (left + right + 1) >> 1
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		}
	}
	dst := left
	left, right, mid = 0, len(matrix[0]), 0
	for left < right {
		mid = (left + right) >> 1
		if matrix[dst][mid] == target {
			return true
		} else if matrix[dst][mid] > target {
			right = mid - 1
		} else if matrix[dst][mid] < target {
			left = mid + 1
		}
	}
	return false
}
