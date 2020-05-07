package main

import (
	"fmt"
)

var not = []byte{110, 111, 32, 97, 110, 115, 119, 101, 114}

func biggerIsGreater(s []byte) []byte {
	i, j := len(s)-1, len(s)-1
	for i > 0 && s[i-1] >= s[i] {
		i--
	}

	if i <= 0 {
		return not
	}

	for s[j] <= s[i-1] {
		j--
	}

	s[i-1], s[j] = s[j], s[i-1]
	for j = len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	var n int
	var s []byte
	fmt.Scanln(&n)

	for n > 0 {
		fmt.Scanln(&s)
		fmt.Printf("%s\n", biggerIsGreater(s))
		n--
	}
}
