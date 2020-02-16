package main

import "fmt"

func main() {
	grid := [][]int{{3, 2}, {-3, -3}, {-3, -3}, {-3, -3}}
	//grid := [][]int{
	//	{4, 3, 2, -1},
	//	{3, 2, 1, -1},
	//	{1, 1, -1, -2},
	//	{-1, -1, -2, -3},
	//}

	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res, flag := 0, n-1
	for i := 0; i < m; i++ {
		for j := flag; j >= 0; j-- {
			if grid[i][j] < 0 {
				res += m - i
				flag -= 1
			} else {
				flag = j
				break
			}
		}
	}

	return res
}
