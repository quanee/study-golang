package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	a string
	b int
}

func main() {
	t := Test{"hello", 2}
	tp := unsafe.Pointer(&t)
	tap := (*string)(unsafe.Pointer(tp))
	*tap = "world"

	tbp := (*int64)(unsafe.Pointer(uintptr(tp) + unsafe.Offsetof(t.b)))
	*tbp = 1
	fmt.Println(t)
}
