package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s but wanted %s", got, want)
		}

	}

	t.Run("should the 'Hello Chris' given a name 'Chris'", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return the 'hello World' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "english")
		want := "hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should be able to handle response in french", func(t *testing.T) {
		got := Hello("William", "french")
		want := "Bonjour William"
		assertCorrectMessage(t, got, want)
	})
	t.Run("should be able to handle a response in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Holla Elodie"
		assertCorrectMessage(t, got, want)
	})
}
