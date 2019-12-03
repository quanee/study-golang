package main

/*
#include <stdint.h>

int64_t myadd(int64_t a, int64_t b) {
	return a+b;
}
*/
import "C"

import (
	"runtime"
	asmpkg "study-golang/asm/07hack/asmcalladd"
	"unsafe"
)

func main() {
	if runtime.GOOS != "windows" {
		println(asmpkg.AsmCallCAdd(
			uintptr(unsafe.Pointer(C.myadd)),
			123, 456,
		))
	}
}
