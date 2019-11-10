package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	got := reconstructMatrix(2, 1, []int{1, 1, 1})
	want := [][]int{[]int{1, 1, 0}, []int{0, 0, 1}}
	want2 := [][]int{[]int{0, 1, 1}, []int{1, 0, 0}}
	want3 := [][]int{[]int{1, 0, 1}, []int{0, 1, 0}}
	if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(got, want2) && !reflect.DeepEqual(got, want3) {
		t.Errorf("Se esperaba %v, %v ó %v pero se recibió %v", want, want2, want3, got)
	}
}

func TestTwo(t *testing.T) {
	got := reconstructMatrix(2, 3, []int{2, 2, 1, 1})
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	got := reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1})
	want := [][]int{[]int{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, []int{1, 0, 1, 0, 0, 0, 1, 1, 0, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestFour(t *testing.T) {
	got := reconstructMatrix(4, 7, []int{2, 1, 2, 2, 1, 1, 1})
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}
