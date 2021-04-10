package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	asserStrings(t, got, want)
}

func asserStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}
