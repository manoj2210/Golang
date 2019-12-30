package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Iterate test function", func(t *testing.T) {
		got := Iterate("M")
		want := "MMMMMMMM"
		assertCorrectMessage(t, got, want)
	})

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a")
	}
}
