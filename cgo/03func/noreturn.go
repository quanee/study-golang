package main

//static void noreturn() {}
import "C"
import "fmt"

func main() {
	v, _ := C.noreturn()
	fmt.Printf("%#v", v)
	fmt.Printf("%v", v)
}
