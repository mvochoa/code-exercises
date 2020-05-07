package main

import "testing"

func TestOne(t *testing.T) {
	got := timeInWords(5, 47)
	want := "thirteen minutes to six"
	if got != want {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := timeInWords(3, 00)
	want := "three o' clock"
	if got != want {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestThree(t *testing.T) {
	got := timeInWords(7, 15)
	want := "quarter past seven"
	if got != want {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestFour(t *testing.T) {
	got := timeInWords(1, 1)
	want := "one minute past one"
	if got != want {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}
