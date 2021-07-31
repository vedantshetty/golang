package main

import "testing"

func TestHello(t *testing.T) {

	assetCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello To People", func(t *testing.T) {
		got := Hello("Vedant", "")
		want := "Hello, Vedant!"

		assetCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assetCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Vedant", "Hindi")
		want := "नमस्ते, Vedant!"
		assetCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Vedant", "French")
		want := "Bonjour, Vedant!"
		assetCorrectMessage(t, got, want)
	})
}
