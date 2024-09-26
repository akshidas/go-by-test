package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is test"
		assertString(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("test_test")
		assertError(t, err)
		assertString(t, err.Error(), ErrUnknownKey.Error())
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add a value", func(t *testing.T) {
		dictionary.Add("test", "This is test")
		got, _ := dictionary.Search("test")
		assertString(t, got, "This is test")
	})

	t.Run("Failed to  add a duplicate value", func(t *testing.T) {
		err := dictionary.Add("test", "This is test")
		assertError(t, err)
		assertString(t, err.Error(), ErrDuplicateKey.Error())
	})

}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	t.Run("update a known key", func(t *testing.T) {
		dictionary.Update("test", "this is test updated")
		got, _ := dictionary.Search("test")
		want := "this is test updated"
		assertString(t, want, got)
	})

	t.Run("update unknown key", func(t *testing.T) {
		got := dictionary.Update("ttest", "this is test updated")
		want := ErrFailedToUpdate
		assertError(t, got)
		assertString(t, want.Error(), got.Error())
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is test"}
	dictionary.Delete("test")
	_, got := dictionary.Search("test")
	want := ErrUnknownKey
	assertError(t, got)
	assertString(t, want.Error(), got.Error())
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error %s but didn't get one", ErrUnknownKey.Error())
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s but want %s", got, want)
	}
}
