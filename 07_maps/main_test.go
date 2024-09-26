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

func assertString(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %s but got %s", want, got)
	}
}

// func TestAdd(t *testing.T) {
// 	d := Dictionary{}
//
// 	err := d.Add("test", "this is test")
// }
