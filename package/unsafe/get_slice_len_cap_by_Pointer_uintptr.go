package main

import (
	"fmt"
	"unsafe"
)

func main() {
	mp := make(map[string]int)
	mp["hello"] = 1
	mp["world"] = 2
	fmt.Println("len:", **(**int)(unsafe.Pointer(&mp)))
}
