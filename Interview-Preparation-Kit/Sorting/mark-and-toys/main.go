package main

import (
	"fmt"
	"sort"
)

func maximumToys(prices []int, k int) int {
	toys := 0

	sort.Ints(prices)
	for _, v := range prices {
		if (k - v) >= 0 {
			k -= v
			toys++
		}
	}

	return toys
}

func main() {
	var s, m int
	fmt.Scanf("%d %d\n", &s, &m)

	a := make([]int, s)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(maximumToys(a, m))
}
