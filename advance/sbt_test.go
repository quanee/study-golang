package main

import (
	"testing"
)

func TestSortFloat64FastV1(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	SortFloat64FastV1(a)
}

func TestSortFloat64FastV2(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	SortFloat64FastV2(a)
}

func BenchmarkSortFloat64FastV1(b *testing.B) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

	for i := 0; i < b.N; i++ {
		SortFloat64FastV1(a)
		//sort.Float64s(a)
	}
}
