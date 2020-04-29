package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search existing word", func(t *testing.T) {
		dictionary := Dictionary(map[string]string{"test": "this is just a test"})

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)

	})
	t.Run("Known word", func(t *testing.T) {
		dictonary := Dictionary{"go": "go there", "find": "seek something"}
		got, err := dictonary.Search("finds")
		if err == nil {
			t.Fatal("error cannot be nil")
		}
		assertStrings(t, got, "")
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("Update word", func(t *testing.T) {
		want := "word for tomorrow"
		dictionary := Dictionary{"word": "word of the day"}
		_, err := dictionary.Update("word", "word for tomorrow")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "word", want)

	})

	t.Run("Add new word", func(t *testing.T) {
		want := "nice place to be"
		definition := "nice"
		dictionary := Dictionary{"go": "go there", "find": "seek something"}
		_, err := dictionary.Add(definition, want)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, definition, want)
	})

	t.Run("Existing word should fail when adding it again", func(t *testing.T) {
		dictionary := Dictionary{"good": "nice"}
		_, err := dictionary.Add("good", "new good")
		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, "good", "nice")
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, want)

}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' instead '%s'", got, want)
	}
}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' instead '%s'", got, want)
	}
}
