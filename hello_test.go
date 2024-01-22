package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Asher")
	want := "Hello, Asher"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
