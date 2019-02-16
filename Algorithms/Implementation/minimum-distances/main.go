package main

import (
	"fmt"
	"sort"
)

func minimumDistances(a []int) int {
	var mins []int

	m := make(map[int][]int)
	for i, v := range a {
		m[v] = append(m[v], i)
	}

	for _, v := range m {
		if len(v) > 1 {
			for i := len(v) - 1; i > 0; i-- {
				mins = append(mins, (v[i] - v[i-1]))
			}
		}
	}

	if len(mins) == 0 {
		return -1
	}

	sort.Ints(mins)
	return mins[0]
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(minimumDistances(a))
}
