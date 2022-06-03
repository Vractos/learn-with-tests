package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestHelloTDD(t *testing.T) {
	got := HelloTDD("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
