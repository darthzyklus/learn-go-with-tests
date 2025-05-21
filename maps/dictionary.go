package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value := d[word]

	if value == "" {
		return "", errors.New("could not find the word you were looking for")
	}

	return value, nil
}
