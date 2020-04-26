package main

import "testing"

func TestOne(t *testing.T) {
	got := workbook(5, 3, []int{4, 2, 6, 1, 10})
	want := 4
	if got != want {
		t.Errorf("workbook(5, 3, []int{4, 2, 6, 1, 10}) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := workbook(10, 5, []int{3, 8, 15, 11, 14, 1, 9, 2, 24, 31})
	want := 8
	if got != want {
		t.Errorf("workbook(10, 5, []int{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}) was incorrect, got: %v, want: %v", got, want)
	}
}
