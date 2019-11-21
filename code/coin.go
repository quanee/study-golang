package main

import (
	"fmt"
)

func coin(deno []int, v int) int {
	n := len(deno)
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, v+1))
		dp[i][0] = 1
	}
	for j := 0; j <= v; j++ {
		if j%deno[0] == 0 {
			dp[0][j] = 1
		}
		/*else {
			dp[0][j] = 0
		}*/
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= v; j++ {
			tmp := 0
			for k := 0; k*deno[i] <= j; k++ {
				tmp += dp[i-1][j-k*deno[i]]
			}
			dp[i][j] = tmp
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[n-1][v]
}

func main() {
	deno := make([]int, 0, 5)
	deno = append(deno, 50, 20, 10, 5, 1)

	/*
		for i := range deno {
			fmt.Println(deno[i])
		}
	*/
	head := make([]int, 0, 101)
	for i := 0; i < 101; i++ {
		head = append(head, i)
	}
	fmt.Println(head)
	c := coin(deno, 100)
	fmt.Println(c)
}
