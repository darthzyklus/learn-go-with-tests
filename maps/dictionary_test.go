package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, err, want)

	})

	t.Run("add new word", func(t *testing.T) {
		dictionary.Add("manual", "set of instructions to explain how something works")

		got, _ := dictionary.Search("manual")
		want := "set of instructions to explain how something works"

		assertStrings(t, got, want)
	})
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected to get an error...")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
