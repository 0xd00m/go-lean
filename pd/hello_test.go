package main

import "testing"

func TestHello(t *testing.T) {
	error_message := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("hello with names", func(t *testing.T) {
		got := hello("chris", "")
		want := "Hello, chris"

		error_message(t, got, want)

	})

	t.Run("hello without names", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world"
		error_message(t, got, want)
	})

	t.Run("hello with spanish", func(t *testing.T) {
		got := hello("emmanuel", "Spanish")
		want := "Hola, emmanuel"
		error_message(t, got, want)
	})
}
