package main

import "fmt"

func main() {
	x := 1
	obj := Constructor1()
	obj.Push(x)
	obj.Push(x)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}

type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{stack: make([]int, 0, 8)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	i := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return i
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
