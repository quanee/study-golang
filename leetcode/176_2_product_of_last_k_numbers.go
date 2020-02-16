package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(3) // [3]
	obj.Add(0) // [3,0]
	obj.Add(2) // [3,0,2]
	obj.Add(5) // [3,0,2,5]
	obj.Add(4) // [3,0,2,5,4]
	fmt.Println(obj.List)
	obj.GetProduct(2) // 返回 20 。最后 2 个数字的乘积是 5 * 4 = 20
	obj.GetProduct(3) // 返回 40 。最后 3 个数字的乘积是 2 * 5 * 4 = 40
	obj.GetProduct(4) // 返回  0 。最后 4 个数字的乘积是 0 * 2 * 5 * 4 = 0
	obj.Add(8)        // [3,0,2,5,4,8]
	obj.GetProduct(2) // 返回 32 。最后 2 个数字的乘积是 4 * 8 = 32
	fmt.Println(obj.GetProduct(3))
}

type ProductOfNumbers struct {
	List []int
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{List: make([]int, 0, 8)}
}

func (this *ProductOfNumbers) Add(num int) {
	this.List = append(this.List, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res, tmp := 0, 0
	for _, v := range this.List[len(this.List)-k:] {
		tmp = v
		fmt.Println(v)
		res *= tmp
	}

	return res
}
