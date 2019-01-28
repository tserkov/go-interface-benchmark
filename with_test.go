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

// interface; non-pointer
func BenchmarkWithFunc(b *testing.B) {
	var t I = &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.F()
	}
}

// interface; pointer
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
		var c T = *t.(*T)

		// Stop the go compiler from being smart.
		_ = c
		_ = t
	}
}

func BenchmarkWithPtrCastAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t I = &T{}
		var c *T = t.(*T)

		// Stop the go compiler from being smart.
		_ = c
		_ = t
	}
}

// struct, cast from interface; non-pointer
func BenchmarkWithCastFunc(b *testing.B) {
	var t I = &T{}
	var c T = *t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.F()
	}
}

// struct, cast from interface; pointer
func BenchmarkWithCastPtrFunc(b *testing.B) {
	var t I = &T{}
	var c T = *t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.PtrF()
	}
}

// *struct, cast from interface; non-pointer
func BenchmarkWithPtrCastFunc(b *testing.B) {
	var t I = &T{}
	var c *T = t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.F()
	}
}

// *struct, cast from interface; pointer
func BenchmarkWithPtrCastPtrFunc(b *testing.B) {
	var t I = &T{}
	var c *T = t.(*T)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.PtrF()
	}
}
