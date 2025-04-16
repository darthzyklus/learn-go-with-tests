package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Andres")
	want := "Hello, Andres"

	if got != want {
		t.Errorf("got: %q; want %q", got, want)
	}
}
