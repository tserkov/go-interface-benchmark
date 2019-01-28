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

// *struct; non-pointer
func BenchmarkWithoutPtrFunc(b *testing.B) {
	var t *T = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

// *struct; pointer
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

// struct; non-pointer
func BenchmarkWithoutNonPtrFunc(b *testing.B) {
	var t T = T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

// struct; pointer
func BenchmarkWithoutNonPtrPtrFunc(b *testing.B) {
	var t T = T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.PtrF()
	}
}
