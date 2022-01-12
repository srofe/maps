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

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
