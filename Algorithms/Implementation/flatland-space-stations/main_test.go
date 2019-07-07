package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	got := flatlandSpaceStations(5, []int{0, 4})
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("flatlandSpaceStations(5, []int{0, 4}) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := flatlandSpaceStations(6, []int{0, 1, 2, 4, 3, 5})
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Errorf("flatlandSpaceStations(6, []int{0, 1, 2, 4, 3, 5}) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestThree(t *testing.T) {
	got := flatlandSpaceStations(20, []int{13, 1, 11, 10, 6})
	want := 6
	if !reflect.DeepEqual(got, want) {
		t.Errorf("flatlandSpaceStations(20, []int{13, 1, 11, 10, 6}) was incorrect, got: %v, want: %v", got, want)
	}
}
