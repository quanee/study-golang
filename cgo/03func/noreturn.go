package main

//static void noreturn() {}
import "C"
import "fmt"

func main() {
	v, _ := C.noreturn()
	fmt.Printf("%#v\n", v)
	fmt.Printf("%v\n", v)
}
