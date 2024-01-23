package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to poeple", func(t *testing.T) {
		got := hello("Elodie", "")
		want := "Hello, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Korean", func(t *testing.T) {
		got := hello("Elodie", "Korean")
		want := "Annyeonghaseyo, Elodie"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
