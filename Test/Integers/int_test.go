package main

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Adding Two Integers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, strconv.Itoa(got), strconv.Itoa(want))
	})

}
