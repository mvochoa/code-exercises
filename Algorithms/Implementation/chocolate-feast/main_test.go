package main

import "testing"

func TestOne(t *testing.T) {
	got := chocolateFeast(10, 2, 5)
	want := 6
	if got != want {
		t.Errorf("chocolateFeast(10, 2, 5) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := chocolateFeast(12, 4, 4)
	want := 3
	if got != want {
		t.Errorf("chocolateFeast(10, 2, 5) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestThree(t *testing.T) {
	got := chocolateFeast(6, 2, 2)
	want := 5
	if got != want {
		t.Errorf("chocolateFeast(10, 2, 5) was incorrect, got: %v, want: %v", got, want)
	}
}
