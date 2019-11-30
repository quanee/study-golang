package main

/*
struct A {
	int   size: 10; // 位字段无法访问
	float arr[];    // 零长的数组也无法访问
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a.size) // 错误: 位字段无法访问
	fmt.Println(a.arr)  // 错误: 零长的数组也无法访问
}
