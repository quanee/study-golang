package main

//void SayHello(char* s);
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
