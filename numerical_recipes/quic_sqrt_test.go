package numrecipes

import "testing"

func TestQRSqrt(t *testing.T) {
	t.Log(QRSqrtF64(16))
	t.Log(QRSqrtF32(16))
	t.Log(Sqrt(16))
}


func BenchmarkQRSqrt(b *testing.B) {
	for i := 0; i < b.N; i++{
		QRSqrtF64(16)
		QRSqrtF32(16)
		Sqrt(16)
	}
}
