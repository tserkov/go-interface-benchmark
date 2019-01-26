package main

import (
	"testing"
)

func BenchmarkWithoutPtrAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t *T = &T{}

		// Stop the go compiler from being smart.
		_ = t
	}
}

func BenchmarkWithoutPtrFunc(b *testing.B) {
	var t *T = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

func BenchmarkWithoutPtrPtrFunc(b *testing.B) {
	var t *T = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.PtrF()
	}
}

func BenchmarkWithoutNonPtrAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t T = T{}

		// Stop the go compiler from being smart.
		_ = t
	}
}

func BenchmarkWithoutNonPtrFunc(b *testing.B) {
	var t T = T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

func BenchmarkWithoutNonPtrPtrFunc(b *testing.B) {
	var t T = T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.PtrF()
	}
}
