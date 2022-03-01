package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people in hindi", func(t *testing.T) {
		got := Hello("Chris", "Hindi")
		want := "Namaste Chris"
		assert(got, want, t)
	})
	t.Run("say hello to people in english", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"
		assert(got, want, t)
	})
	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola Chris"
		assert(got, want, t)
	})
	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour Chris"
		assert(got, want, t)
	})
	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"
		assert(got, want, t)
	})
}

func assert(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
