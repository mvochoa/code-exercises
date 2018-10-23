package main

import (
	"fmt"
)

func equalizeArray(arr []int) int {
	t, max, index := 0, 1, arr[0]
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
		if m[v] > max {
			max = m[v]
			index = v
		}
	}

	for _, v := range arr {
		if index != v {
			t++
		}
	}

	return t
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(equalizeArray(a))
}
