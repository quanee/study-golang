package numrecipes

import "testing"

func TestGCD(t *testing.T) {
	t.Log(GCD(15, 2))
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(uint(i), uint(i+1))
	}
}
