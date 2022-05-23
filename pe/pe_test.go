package pe

import (
	"testing"
)

func TestSumMultiples(t *testing.T) {
	want := 233168
	got := SumMultiples(1000, []int{3, 5})
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestSumEvenFibonacci(t *testing.T) {
	want := 4613732
	got := SumEvenFibonacci(4000000)
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	want := 6857
	got := LargestPrimeFactor(600851475143)
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
