package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello with names", func(t *testing.T) {
		got := hello("chris")
		want := "Hello, chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("hello without names", func(t *testing.T) {
		got := hello()
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
