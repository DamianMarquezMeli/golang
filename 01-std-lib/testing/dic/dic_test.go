package dictionary_test

import (
	"testing"

	dic "github.com/devpablocristo/go-concepts/std-lib/testing/dic"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test", "hello": "hi"}

	got := dic.Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
