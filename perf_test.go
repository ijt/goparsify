package goparsify

import "testing"

func BenchmarkAny(b *testing.B) {
	p := Any("hello", "goodbye", "help")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = Run(p, "hello")
		_, _, _ = Run(p, "hello world")
		_, _, _ = Run(p, "good boy")
		_, _, _ = Run(p, "help me")
	}
}
