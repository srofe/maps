package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("A Dictionary shall have a Search method to find elements it contains.", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("The Search method shall return an error if the element is not contained in the map.", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you are looking for"
		if err == nil {
			t.Fatal("Expected to get an error.")
		}
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("A Dictionary shall allow words and definitions to be added.", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find word: ", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
