package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		err, got := dictionary.Search("test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		err, _ := dictionary.Search("unknown")

		assertError(t, err, NotFoundError)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, WordExistsError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		newDefinition := "new test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		assertError(t, err, WordDoesNotExistError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		err, _ := dictionary.Search(word)

		if err != NotFoundError {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

func assertDefinition(t testing.TB, d Dictionary, word string, definition string) {
	err, got := d.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
