package main

import "testing"

func TestOne(t *testing.T) {
	got := fairRations([]int{2, 3, 4, 5, 6})
	want := 4
	if got != want {
		t.Errorf("fairRations([]int{2, 3, 4, 5, 6}]) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := fairRations([]int{1, 2})
	want := -1
	if got != want {
		t.Errorf("fairRations([]int{1, 2}) was incorrect, got: %v, want: %v", got, want)
	}
}
