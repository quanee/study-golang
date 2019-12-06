package main

func copyslice(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func appendslice(a []int) []int {
	b := append(a[:0:0], a...)
	return b
}
