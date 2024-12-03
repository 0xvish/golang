package main

import "testing"

// TestHello tests the Hello function
func TestHello(t *testing.T) {

	// Subtest
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Vish⚡")
		want := "Hello, Vish⚡!"

		if got != want {
			assertCorrectMesssage(t, got, want)
		}
	})

	// Subtest
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got:= Hello("")
		want:= "Hello, World!"

		if got != want {
			assertCorrectMesssage(t, got, want)
		}
	})
}

// assertCorrectMesssage is a helper function to reduce duplication in our tests
func assertCorrectMesssage(t testing.TB, got, want string) { // t testing.TB is an interface that *testing.T satisfies. This means we can call our helper function from any object that satisfies the testing.TB interface. This includes *testing.T and *testing.B
	t.Helper() // This tells the test suite that the method is a helper. When it fails, the line number reported will be in our function call rather than inside our test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}