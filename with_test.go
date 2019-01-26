package main

import (
	"testing"
)

func BenchmarkWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t I = &T{}

		// Stop the go compiler from being smart.
		_ = t
	}
}

func BenchmarkWithFunc(b *testing.B) {
	var t I = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

func BenchmarkWithPtrFunc(b *testing.B) {
	var t I = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.PtrF()
	}
}

func BenchmarkWithCastAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t I = &T{}
		var c *T = t.(*T)

		// Stop the go compiler from being smart.
		_ = c
		_ = t
	}
}

func BenchmarkWithCastFunc(b *testing.B) {
	var t I = &T{}
	var c *T = t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.F()
	}
}

func BenchmarkWithCastPtrFunc(b *testing.B) {
	var t I = &T{}
	var c *T = t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.PtrF()
	}
}
