package main

import (
	"fmt"
)

func pickingNumbers(a []int) int {
	var n int
	max := make([]int, len(a))
	for i := range a {
		for j := range a {
			if a[i] == a[j] || a[i]-1 == a[j] {
				max[i]++
			}
		}
	}

	for _, v := range max {
		if v > n {
			n = v
		}
	}

	return n
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	fmt.Println(pickingNumbers(arr))
}
