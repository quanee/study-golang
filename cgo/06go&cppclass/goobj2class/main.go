package main

// #include "./person_capi.h"
// #cgo CXXFLAGS: -std=c++11
// extern void Main();
import "C"

func main() {
	C.Main()
}
