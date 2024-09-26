package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is test"}

	t.Run("searching known key", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "this is test"
		assertString(t, want, got)
	})

	t.Run("searching unknown key", func(t *testing.T) {
		_, got := d.Search("ttest")
		want := ErrUnknownKey
		assertError(t, want, got)
		assertString(t, want.Error(), got.Error())
	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{}

	t.Run("Add a new key", func(t *testing.T) {
		got := d.Add("test", "this is test")

		if got != nil {
			t.Error("Wanted Error to be nil but got an error")
		}
	})

	t.Run("Add a duplicate key", func(t *testing.T) {
		got := d.Add("test", "this is test")
		want := ErrDuplicateKey

		assertError(t, want, got)
		assertString(t, want.Error(), got.Error())
	})

}

func assertError(t testing.TB, want, got error) {
	t.Helper()
	if got == nil {
		t.Fatalf("Want error to be %v but got %v", want, got)
	}
}

func assertString(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s but got %s", want, got)
	}
}
