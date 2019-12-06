package main

import "testing"

func BenchmarkStringJoinOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fc()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fd()
	}
}
