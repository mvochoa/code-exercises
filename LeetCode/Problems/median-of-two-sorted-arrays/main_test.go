package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	got := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	want := float64(2.5)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestTwo(t *testing.T) {
	got := findMedianSortedArrays([]int{1, 3}, []int{2})
	want := float64(2)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	got := findMedianSortedArrays([]int{}, []int{1})
	want := float64(1)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFour(t *testing.T) {
	got := findMedianSortedArrays([]int{3}, []int{-2, -1})
	want := float64(-1)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFive(t *testing.T) {
	got := findMedianSortedArrays([]int{1, 2}, []int{-1, 3})
	want := float64(1.5)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestSix(t *testing.T) {
	got := findMedianSortedArrays([]int{1, 2}, []int{1, 2, 3})
	want := float64(2)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestSeven(t *testing.T) {
	got := findMedianSortedArrays([]int{1, 3}, []int{2, 4})
	want := float64(2.5)
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}
