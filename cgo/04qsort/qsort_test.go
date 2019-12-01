package qsort

import (
	_sort "sort"
	"testing"
)

func TestSort(t *testing.T) {
	values := []int32{42, 9, 101, 95, 27, 25}

	Sort(values)

	isSorted := _sort.SliceIsSorted(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	if !isSorted {
		t.Fatal("should be sorted")
	}
}
