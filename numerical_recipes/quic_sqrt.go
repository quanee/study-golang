package numrecipes

import (
	"unsafe"
)

func QRSqrtF64(number float64) float64 {
	var i int
	var twohalfs, y float64
	const threehalfs float64 = 1.5

	twohalfs = number * 0.5
	y = number
	i = *(*int)(unsafe.Pointer(&y))
	i = 0x5fe6ec85e7de30da - (i >> 1)
	y = *(*float64)(unsafe.Pointer(&i))
	y = y * (threehalfs - (twohalfs * y * y))

	return y
}

func QRSqrtF32(number float32) float32 {
	var i int
	var x2, y float32
	const threehalfs float32 = 1.5

	x2 = number * 0.5
	y = number
	i = *(*int)(unsafe.Pointer(&y))
	i = 0x5f3759df - (i >> 1) // John Carmack‘s number
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - (x2 * y * y))
	//y = y * (threehalfs - (x2 * y * y))

	return y
}

// Newton's Method
func Sqrt(number float64) float64 {
	if number == 0 {
		return 0
	}
	last, res := 0.0, 1.0
	for res != last {
		last = res
		res = (res + number/res) / 2
	}
	return res
}
