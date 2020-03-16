package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func stringtobyte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	str := "hello"
	sli := []byte("world")

	fmt.Printf("%[1]v, %[1]T", stringtobyte(str))
	fmt.Printf("%[1]v, %[1]T", bytes2string(sli))
}
