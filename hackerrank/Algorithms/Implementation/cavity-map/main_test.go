package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	got := cavityMap([]string{"1112", "1912", "1892", "1234"})
	want := []string{"1112", "1X12", "18X2", "1234"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestTwo(t *testing.T) {
	got := cavityMap([]string{"989", "191", "111"})
	want := []string{"989", "1X1", "111"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}

func TestThree(t *testing.T) {
	got := cavityMap([]string{"12", "12"})
	want := []string{"12", "12"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("timeInWords(5, 47) was incorrect, got: %v, want: %v", got, want)
	}
}
