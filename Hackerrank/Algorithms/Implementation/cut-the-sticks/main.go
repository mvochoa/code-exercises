package main

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int) []int {
	var res []int
	l, pos := len(arr), 0

	sort.Ints(arr)
	for i := 0; i < l-1; i++ {
		if arr[i] != arr[i+1] {
			res = append(res, (l - pos))
			pos = i + 1
		}
	}

	return append(res, (l - pos))
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	for _, v := range cutTheSticks(arr) {
		fmt.Println(v)
	}
}
