package numrecipes

import "testing"

func TestQPow(t *testing.T) {
	t.Log(QPow(300, 300))
}

func BenchmarkQPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QPow(300, i)
	}
}
