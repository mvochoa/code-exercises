package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	got := stones(3, 2, 3)
	want := []int{4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("stones(3, 2, 3) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := stones(3, 1, 2)
	want := []int{2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("stones(3, 1, 2) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestThree(t *testing.T) {
	got := stones(4, 10, 100)
	want := []int{30, 120, 210, 300}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("stones(4, 10, 100) was incorrect, got: %v, want: %v", got, want)
	}
}
