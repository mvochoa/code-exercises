package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	width := []int{2, 3, 1, 2, 3, 2, 3, 3}
	cases := [][]int{
		[]int{0, 3},
		[]int{4, 6},
		[]int{6, 7},
		[]int{3, 5},
		[]int{0, 7},
	}
	got := serviceLane(width, cases)
	want := []int{1, 2, 3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("serviceLane({2, 3, 1, 2, 3, 2, 3, 3}, { {0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	width := []int{1, 2, 2, 2, 1}
	cases := [][]int{
		[]int{2, 3},
		[]int{1, 4},
		[]int{2, 4},
		[]int{2, 4},
		[]int{2, 3},
	}
	got := serviceLane(width, cases)
	want := []int{2, 1, 1, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("serviceLane({1, 2, 2, 2, 1}, { {2, 3}, {1, 4}, {2, 4}, {2, 4}, {2, 3} }) was incorrect, got: %v, want: %v", got, want)
	}
}
