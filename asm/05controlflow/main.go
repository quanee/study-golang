package main

import "fmt"

func If(ok int, a, b int) int
func seq()
func LoopAdd(cnt, v0, step int) int

func main() {
	//seq()
	fmt.Println(If(0, 1, 2))
	fmt.Println(If(1, 1, 2))
	fmt.Println(LoopAdd(1, 2, 3))
}
