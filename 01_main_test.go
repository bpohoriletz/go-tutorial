package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Logf("got %q, want %q", got, want)
		t.Fail()
	}
}

func Test_Hello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello, World!' when empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}
