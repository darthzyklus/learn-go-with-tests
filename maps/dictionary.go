package main

type Dictionary map[string]string

const (
	NotFoundError         = DictionaryErr("could not find the word you were looking for")
	WordExistsError       = DictionaryErr("cannot add word because it already exists")
	WordDoesNotExistError = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (error, string) {
	definition, ok := d[word]

	if !ok {
		return NotFoundError, ""
	}

	return nil, definition
}

func (d Dictionary) Add(word string, definition string) error {
	err, _ := d.Search(word)

	if err == nil {
		return WordExistsError
	}

	if err == NotFoundError {
		d[word] = definition
		return nil
	}

	return err
}

func (d Dictionary) Update(word string, definition string) error {
	err, _ := d.Search(word)

	if err == NotFoundError {
		return WordDoesNotExistError
	}

	if err == nil {
		d[word] = definition
	}

	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
