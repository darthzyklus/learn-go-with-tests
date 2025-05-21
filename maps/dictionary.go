package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if ok {
		return value, nil
	}

	return "", errors.New("could not find the word you were looking for")
}
