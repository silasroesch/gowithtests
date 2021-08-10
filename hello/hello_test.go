package main

import "testing"

func TestHello(t *testing.T) {
	assertCorretMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to Yakub", func(t *testing.T) {
		got := Hello("Yakub", "")
		want := "Hello, Yakub"
		assertCorretMessage(t, got, want)
	})

	t.Run("saying Hello to Max", func(t *testing.T) {
		got := Hello("Max", "")
		want := "Hello, Max"
		assertCorretMessage(t, got, want)
	})

	t.Run("in german", func(t *testing.T) {
		got := Hello("Conner", "german")
		want := "Hallo, Conner"
		assertCorretMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Conner", "french")
		want := "Bonjour, Conner"
		assertCorretMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Conner", "spanish")
		want := "Hola, Conner"
		assertCorretMessage(t, got, want)
	})

	t.Run("say 'Hello, World' if no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorretMessage(t, got, want)
	})
}
