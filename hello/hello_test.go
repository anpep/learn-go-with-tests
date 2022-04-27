package hello

import "testing"

func TestHello(t *testing.T) {
	assertMessageIsCorrect := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Ángel", "")
		want := "Hello, Ángel"
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello to people in the default language (english)", func(t *testing.T) {
		got := Hello("Ángel", "english")
		want := Hello("Ángel", "")
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Javier", "spanish")
		want := "Hola Javier"
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Benoît", "french")
		want := "Bonjour Benoît"
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello world by default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello world in the default language (english)", func(t *testing.T) {
		got := Hello("", "english")
		want := Hello("", "")
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello world in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola mundo"
		assertMessageIsCorrect(t, got, want)
	})

	t.Run("say hello world in french", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour monde"
		assertMessageIsCorrect(t, got, want)
	})

}
