package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "a test is a test"}

	t.Run("known word", func(t *testing.T) {
		// DECLARING A MAP
		// map [<keytype>] <value type>
		// key: value => "videogame" : "xbox"
		// The value type, can be anything, even ANOTHER MAP
		// dictionary := map[string]string{"test": "a test is a test"}

		got, _ := dictionary.Search("test")
		want := "a test is a test"

		checkStrings(t, got, want)
	})
	t.Run("UNknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		checkError(t, got, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		checkError(t, err, ErrWordDoesNotExists)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		checkError(t, err, nil)
		checkDefinition(t, dictionary, word, newDefinition)

	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)
	checkDefinition(t, dictionary, word, newDefinition)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition for delete"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word) //check if the word has been removed
	checkError(t, err, ErrNotFound)
}

func checkStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func checkError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func checkDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	checkStrings(t, got, definition)
}
