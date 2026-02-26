package main

import "testing"

func BenchmarkConcatSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatSlow(1000)
	}
}

func BenchmarkConcatFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatFast(1000)
	}
}

// go test -bench=. -benchmem
