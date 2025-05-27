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
		assertError(t, err, ErrNotFound)

	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		dictionary.Add("manual", "set of instructions to explain how something works")

		got, _ := dictionary.Search("manual")
		want := "set of instructions to explain how something works"

		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		ssertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestRemove(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("existing word", func(t *testing.T) {
		dictionary.Add("manual", "set of instructions to explain how something works")

		err := dictionary.Remove("manual")

		if err != nil {
			t.Error("should not get an error")
		}

		_, err = dictionary.Search("manual")

		assertError(t, err, ErrNotFound)
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
