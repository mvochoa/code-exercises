package main

import "testing"

func TestOne(t *testing.T) {
	got := closedIsland([][]int{[]int{1, 1, 1, 1, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0, 1, 1, 0}, []int{1, 0, 1, 0, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0, 1, 0, 1}, []int{1, 1, 1, 1, 1, 1, 1, 0}})
	want := 2
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestTwo(t *testing.T) {
	got := closedIsland([][]int{[]int{0, 0, 1, 0, 0}, []int{0, 1, 0, 1, 0}, []int{0, 1, 1, 1, 0}})
	want := 1
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	got := closedIsland([][]int{[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 0, 0, 0, 0, 0, 1},
		[]int{1, 0, 1, 1, 1, 0, 1},
		[]int{1, 0, 1, 0, 1, 0, 1},
		[]int{1, 0, 1, 1, 1, 0, 1},
		[]int{1, 0, 0, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1}})
	want := 2
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFour(t *testing.T) {
	got := closedIsland([][]int{[]int{0, 0, 1, 1, 0, 1, 0, 0, 1, 0}, []int{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}, []int{1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, []int{0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 0}, []int{0, 1, 0, 1, 0, 1, 0, 1, 1, 1}, []int{1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, []int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0}, []int{1, 1, 1, 0, 0, 1, 0, 1, 0, 1}, []int{1, 1, 1, 0, 1, 1, 0, 1, 1, 0}})
	want := 5
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFive(t *testing.T) {
	got := closedIsland([][]int{[]int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1}, []int{0, 0, 1, 0, 0, 1, 0, 1, 1, 1}, []int{1, 0, 1, 0, 0, 0, 1, 0, 1, 0}, []int{1, 1, 1, 1, 1, 0, 0, 1, 0, 0}, []int{1, 0, 1, 0, 1, 1, 1, 1, 1, 0}, []int{0, 0, 0, 0, 1, 1, 0, 0, 0, 0}, []int{1, 0, 1, 0, 0, 0, 0, 1, 1, 0}, []int{1, 1, 0, 0, 1, 1, 0, 0, 0, 0}, []int{0, 0, 0, 1, 1, 0, 1, 1, 1, 0}, []int{1, 1, 0, 1, 0, 1, 0, 0, 1, 0}})
	want := 4
	if got != want {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}
