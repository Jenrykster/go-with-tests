package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("João")
	want := "Hello, João"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
