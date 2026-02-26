package main

import "testing"

// BenchmarkConcatSlow tests naive string concatenation.
// By using +=, Go has to allocate memory, copy the old string, and attach the new char on every single loop!
// Result: 530,277 Bytes of RAM wasted per function call!
func BenchmarkConcatSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatSlow(1000)
	}
}

// BenchmarkConcatFast tests strings.Builder.
// It pre-allocates an internal byte slice and appends directly to it, with zero intermediate copies.
// Result: Only 3,320 Bytes of RAM used, and runs 65x FASTER!
func BenchmarkConcatFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatFast(1000)
	}
}

// go test -bench=. -benchmem
