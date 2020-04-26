package main

import (
	"fmt"
)

func utopianTree(n int) int {
	h := 1
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			h++
		} else {
			h *= 2
		}
	}

	return h
}

func main() {
	var n int
	fmt.Scanln(&n)

	for i := n - 1; i >= 0; i-- {
		fmt.Scanln(&n)
		fmt.Println(utopianTree(n))
	}
}
