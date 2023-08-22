package main 

import "testing"

func TestHello(t *testing.T) {

	// if only name is passed then default is english
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Danish", "")
		want := "Hello, Danish"
		assertCorrectMessage(t, got, want)
	})

	// if no name an no language passed 
	t.Run("saying hello to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	// validate the name in the spanish
	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Pablo","Spanish")
		want := "Hola, Pablo"
		assertCorrectMessage(t, got, want)
	})

	// validate the name in the french 
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Emma","French")
		want := "Bonjour, Emma"
		assertCorrectMessage(t, got, want)
	})

	// validate the name in the hindi 
	t.Run("say hello in Hindi", func(t *testing.T) {
		got := Hello("Kaushik","Hindi")
		want := "Namaste, Kaushik"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}