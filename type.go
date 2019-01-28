package main

type T struct{}

//go:noinline
func (t T) F() {}

//go:noinline
func (t *T) PtrF() {}
