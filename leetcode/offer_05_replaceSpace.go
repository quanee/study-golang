package main

import "fmt"

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	res := make([]rune, 0, len(s)*3)
	for _, si := range s {
		if si == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, si)
		}
	}
	ress := append([]int{3}, 3, 3)

	return string(res)
}
