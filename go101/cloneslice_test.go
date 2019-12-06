package main

import "testing"

var a = make([]int, 100000000)

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyslice(a)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendslice(a)
	}
}
