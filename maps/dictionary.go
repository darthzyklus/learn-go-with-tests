package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("could not add word, it already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		d[word] = definition
		return nil
	}

	if err != nil {
		return err
	}

	return ErrWordExists
}

func (d Dictionary) Remove(word string) error {
	_, ok := d[word]

	if !ok {
		return ErrNotFound
	}

	delete(d, word)

	return nil
}
