package main

import (
	"fmt"
)

func repeatedString(s string, n int) int {
	c := 0
	for i := range s {
		if 'a' == s[i] {
			c++
		}
	}

	c *= int(n / len(s))
	for i := 0; i < (n % len(s)); i++ {
		if 'a' == s[i] {
			c++
		}
	}

	return c
}

func main() {
	var s string
	var n int
	fmt.Scanf("%s\n", &s)
	fmt.Scan(&n)

	fmt.Println(repeatedString(s, n))
}
