package _2ctype

/*
struct A {
	int type;    // type 是 Go 语言的关键字
	float _type; // 将屏蔽CGO对 type 成员的访问
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a._type) // _type 对应 type
}
