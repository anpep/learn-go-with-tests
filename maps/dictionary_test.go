package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStringsEqual(t, got, "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNoSuchWord)
	})

	t.Run("add words", func(t *testing.T) {
		err := dictionary.Add("unknown", "not known")
		assertNoError(t, err)
		got, err := dictionary.Search("unknown")
		assertNoError(t, err)
		assertStringsEqual(t, got, "not known")
		err = dictionary.Add("unknown", "another meaning")
		assertError(t, err, ErrWordAlreadyExists)
	})

	t.Run("update words", func(t *testing.T) {
		err := dictionary.Update("sentinel", "a value")
		assertError(t, err, ErrUpdateNonExistingWord)
		err = dictionary.Update("test", "this is a great test")
		assertNoError(t, err)
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStringsEqual(t, got, "this is a great test")
	})

	t.Run("delete words", func(t *testing.T) {
		_, err := dictionary.Search("test")
		assertNoError(t, err)
		dictionary.Delete("test")
		_, err = dictionary.Search("test")
		assertError(t, err, ErrNoSuchWord)
	})
}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
