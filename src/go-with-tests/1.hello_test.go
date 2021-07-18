package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Vedant")
	want := "Hello, Vedant!"

	if got != want {
		t.Errorf("Hello() = %v, want %v", got, want)
	}
}
