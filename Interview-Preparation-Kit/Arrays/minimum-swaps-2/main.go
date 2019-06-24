package main

import (
	"fmt"
)

func minimumSwaps(a []int) int {
	s, b, aux := 0, 0, 0

	for i := 0; i < len(a); i++ {
		b = a[i] - (i + 1)
		if b < 0 {
			aux = a[i]
			a[i] = a[i+b]
			a[i+b] = aux
			s++
			i--
		}
	}

	return s
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(minimumSwaps(a))
}
