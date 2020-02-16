package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res1 := getAll(root1)
	res2 := append(res1, getAll(root2)...)
	sort(res2, 0, len(res2)-1)
	return res2
}

func sort(res []int, start, end int) {
	if start < end {
		i, j := start, end
		key := res[(start+end)/2]
		for i <= j {
			for res[i] < key {
				i++
			}
			for res[j] > key {
				j--
			}
			if i <= j {
				res[i], res[j] = res[j], res[i]
				i++
				j--
			}
		}

		if start < j {
			sort(res, start, j)
		}
		if end > i {
			sort(res, i, end)
		}
	}
}

func getAll(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			v := stack[len(stack)-1]
			root = v.Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func main() {

}
