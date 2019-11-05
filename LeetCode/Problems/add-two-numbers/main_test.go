package main

import (
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	a := arrayToList([]int{2, 4, 3})
	b := arrayToList([]int{5, 6, 4})
	got := listToArray(addTwoNumbers(a, b))
	want := []int{7, 0, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestTwo(t *testing.T) {
	a := arrayToList([]int{1})
	b := arrayToList([]int{9, 9})
	got := listToArray(addTwoNumbers(a, b))
	want := []int{0, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func TestThree(t *testing.T) {
	a := arrayToList([]int{1, 8})
	b := arrayToList([]int{0})
	got := listToArray(addTwoNumbers(a, b))
	want := []int{1, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v pero se recibió %v", want, got)
	}
}

func listToArray(l *ListNode) []int {
	r := []int{l.Val}
	for l.Next != nil {
		l = l.Next
		r = append(r, l.Val)
	}

	return r
}

func arrayToList(r []int) *ListNode {
	o := &ListNode{
		Val: r[0],
	}
	l := o

	for i := 1; i < len(r); i++ {
		l.Next = &ListNode{
			Val: r[i],
		}
		l = l.Next
	}

	return o
}
