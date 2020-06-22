package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	assertString := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "word not found"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), want)

	})
}

func TestAdd(t *testing.T) {
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertDefinition := func(t *testing.T, dictionary Dictionary, word, definition string) {
		t.Helper()

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if definition != got {
			t.Errorf("got %q want %q", got, definition)
		}
	}

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
