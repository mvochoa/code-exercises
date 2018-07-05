package main

import (
	"fmt"
)

func circularArrayRotation(a []int, k int, queries []int) []int {
	var p, l int = 0, len(a)
	for i, v := range queries {
		p = v - k
		for p < 0 {
			p += l
		}
		queries[i] = a[p]
	}
	return queries
}

func main() {
	var n, k, q int
	fmt.Scanf("%d %d %d\n", &n, &k, &q)

	arr := make([]int, n)
	queries := make([]int, q)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for i := range queries {
		fmt.Scanln(&queries[i])
	}
	for _, v := range circularArrayRotation(arr, k, queries) {
		fmt.Println(v)
	}
}
