package main

import (
	"fmt"
)

func rotLeft(a []int, d int) []int {
	return append(a[d:], a[:d]...)
}

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	r := rotLeft(a, d)

	fmt.Print(r[0])
	for i := 1; i < len(r); i++ {
		fmt.Printf(" %d", r[i])
	}
}
