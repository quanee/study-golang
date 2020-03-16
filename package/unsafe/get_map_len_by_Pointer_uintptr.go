package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sl := make([]int, 4, 10)
	fmt.Println("len:", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sl)) + uintptr(8))))
	fmt.Println("cap:", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sl)) + uintptr(16))))
}
