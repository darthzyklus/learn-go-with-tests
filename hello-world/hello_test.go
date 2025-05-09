package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andres", "")
		want := "Hello, Andres"

		assertEqual(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertEqual(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Andres", "Spanish")
		want := "Hola, Andres"

		assertEqual(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Andres", "French")
		want := "Bonjour, Andres"

		assertEqual(t, got, want)
	})
}

func assertEqual[T comparable](t testing.TB, got T, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want %v", got, want)
	}
}
