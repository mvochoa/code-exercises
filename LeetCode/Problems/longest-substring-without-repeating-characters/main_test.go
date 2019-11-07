package main

import (
	"io/ioutil"
	"testing"
)

func TestOne(t *testing.T) {
	got := lengthOfLongestSubstring("abcabcbb")
	want := 3
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestTwo(t *testing.T) {
	got := lengthOfLongestSubstring("bbbbb")
	want := 1
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	got := lengthOfLongestSubstring("pwwkew")
	want := 3
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFour(t *testing.T) {
	got := lengthOfLongestSubstring("ohomm")
	want := 3
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFive(t *testing.T) {
	got := lengthOfLongestSubstring("")
	want := 0
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestSix(t *testing.T) {
	got := lengthOfLongestSubstring(" ")
	want := 1
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestSeven(t *testing.T) {
	got := lengthOfLongestSubstring("au")
	want := 2
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestEight(t *testing.T) {
	got := lengthOfLongestSubstring("anviaj")
	want := 5
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestNine(t *testing.T) {
	b, _ := ioutil.ReadFile("testcase/case1.txt")
	got := lengthOfLongestSubstring(string(b))
	want := 95
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}
