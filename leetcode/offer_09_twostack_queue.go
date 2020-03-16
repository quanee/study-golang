package main

func main() {

}

type CQueue struct {
	tail []int
	head []int
}

func Constructor_() CQueue {
	return CQueue{tail: make([]int, 0), head: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.tail = append(this.tail, value)
}

func (this *CQueue) DeleteHead() int {
	res := -1
	if len(this.head) == 0 {
		if len(this.tail) == 0 {
			return res
		} else {
			this.head = this.tail
			this.tail = []int{}
		}
	}
	res = this.head[0]
	this.head = this.head[1:]

	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
