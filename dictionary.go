package main

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	value, ok := dictionary[word]
	if !ok {
		return "", errors.New("could not find the word you are looking for")
	}
	return value, nil
}
