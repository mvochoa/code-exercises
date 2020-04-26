package main

import (
	"fmt"
)

func getMoneySpent(keyboards []int, drives []int, b int) int {
	max := -1
	for _, k := range keyboards {
		for _, d := range drives {
			if (k+d) >= max && (k+d) <= b {
				max = (k + d)
			}
		}
	}

	return max
}

func main() {
	var b, n, m int
	fmt.Scanf("%d %d %d\n", &b, &n, &m)

	k := make([]int, n)
	d := make([]int, m)
	for i := range k {
		fmt.Scan(&k[i])
	}
	for i := range d {
		fmt.Scan(&d[i])
	}

	fmt.Println(getMoneySpent(k, d, b))
}
