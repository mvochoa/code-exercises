package main

import "testing"

func TestOne(t *testing.T) {
	got := oddCells(2, 3, [][]int{[]int{0, 1}, []int{1, 1}})
	want := 6
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestTwo(t *testing.T) {
	got := oddCells(2, 2, [][]int{[]int{1, 1}, []int{0, 0}})
	want := 0
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	got := oddCells(3, 3, [][]int{[]int{1, 1}, []int{0, 0}, []int{1, 2}})
	want := 6
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFour(t *testing.T) {
	got := oddCells(4, 3, [][]int{[]int{1, 1}})
	want := 5
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}
