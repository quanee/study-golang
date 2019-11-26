package advance

import "testing"

func TestSortFloat64FastV1(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	SortFloat64FastV1(a)
}

func TestSortFloat64FastV2(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	SortFloat64FastV2(a)
}
