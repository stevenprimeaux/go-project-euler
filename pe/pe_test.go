package pe

import (
	"testing"
)

func TestSumMultiples(t *testing.T) {
	want := 23
	got := SumMultiples(10, []int{3, 5})
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}

	want = 233168
	got = SumMultiples(1000, []int{3, 5})
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
