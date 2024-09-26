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
		want := ErrUnknownKey.Error()
		if got == nil {
			t.Fatalf("Wanted error %v but got nil", want)
		}
		assertString(t, want, got.Error())
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
		want := "duplicate key"
		if got == nil {
			t.Fatal("Wanted error to be duplicate key but got nil")
		}

		if got.Error() != want {
			t.Errorf("Wanted %v but got %v", want, got)
		}
	})

}

func assertString(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s but got %s", want, got)
	}
}
