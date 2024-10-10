package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Forest")
  want := "Hello, Forest"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
