package main

//#cgo CPPFLAGS: -I./
//void SayHello(const char* s);
import "C"

func main() {
	C.SayHello(C.CString("Hello World\n"))
}
