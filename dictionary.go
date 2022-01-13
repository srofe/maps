package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you are looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)
	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}
	return nil
}
