package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Nathan", "")
		want := "Hello, Nathan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say language", func(t *testing.T) {
		t.Run("say hello in spanish", func(t *testing.T) {
			got := Hello("Nathan", "Spanish")
			want := "Hola, Nathan"
			assertCorrectMessage(t, got, want)
		})
		t.Run("say hello in french", func(t *testing.T) {
			got := Hello("Nathan", "French")
			want := "Bonjour, Nathan"
			assertCorrectMessage(t, got, want)
		})
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
