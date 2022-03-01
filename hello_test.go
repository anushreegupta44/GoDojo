package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		assert(got, want, t)
	})
	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
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
