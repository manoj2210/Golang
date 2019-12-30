package main

import (
	"testing"
)

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func TestArray(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testing the Array", func(t *testing.T) {
		number := []int{1, 3, 4, 5}
		got := Sum(number)
		want := sum(number)
		assertCorrectMessage(t, got, want)
	})
}
