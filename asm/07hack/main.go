package main

import "runtime"

func SyscallWrite_Darwin(fd int, msg string) int

func main() {
	if runtime.GOOS == "darwin" {
		SyscallWrite_Darwin(1, "hello syscall!\n")
	}
}
