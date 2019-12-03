package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

const g_goid_offset = 152 // Go1.10

//go:nosplit
func getg() unsafe.Pointer

func GetGroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func main() {
	fmt.Println(GetGoid())
	fmt.Println(GetGroutineId())
}
