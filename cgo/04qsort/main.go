package _4qsort

import (
	"fmt"

	qsort "study-golang/cgo/04qsort/qsortV1"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Sort(values)

	fmt.Println(values)
}
