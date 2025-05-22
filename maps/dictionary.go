package main

import "errors"

type Dictionary map[string]string

var translations = map[string]Dictionary{}

var SupportedLanguages = map[string]bool{
	"spanish": true,
}

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrNotSupportedLanguage = errors.New("language is not supported")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}

func (d Dictionary) Remove(word string) error {
	_, ok := d[word]

	if !ok {
		return ErrNotFound
	}

	delete(d, word)

	return nil
}

func (d Dictionary) Translate(word string, lang string) (string, error) {
	_, ok := SupportedLanguages[lang]

	if !ok {
		return "", ErrNotSupportedLanguage
	}

	dic, ok := translations[lang]

	if !ok {
		return "", ErrNotSupportedLanguage
	}

	value, err := dic.Search(word)

	return value, err

}

func (d Dictionary) AddTranslation(word string, translation string, lang string) error {
	_, ok := SupportedLanguages[lang]

	if !ok {
		return ErrNotSupportedLanguage
	}

	dic, ok := translations[lang]

	if !ok {
		translations[lang] = Dictionary{}
		dic = translations[lang]
	}

	dic.Add(word, translation)

	return nil
}
