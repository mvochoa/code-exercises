package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	got := hello("Welcome to 30 Days of Code!")
	want := "Hello, World.\nWelcome to 30 Days of Code!"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba '%v' pero se recibió '%v'.", want, got)
	}
}
