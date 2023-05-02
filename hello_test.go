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

func TestHelloAgain(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := HelloTDD("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty string defaults to 'World'", func(t *testing.T) {
		got := HelloTDD("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
