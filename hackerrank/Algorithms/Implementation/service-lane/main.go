package main

import (
	"fmt"
)

func minIntSlice(v []int) int {
	min := v[0]
	for _, v := range v {
		if v < min {
			min = v
		}
	}
	return min
}

func serviceLane(width []int, cases [][]int) []int {
	result := make([]int, len(cases))
	for i, v := range cases {
		result[i] = minIntSlice(width[v[0]:(v[1]+1)])
	}
	return result
}

func main() {
	var n, c int
	fmt.Scanf("%d %d\n", &n, &c)

	width := make([]int, n)
	for i := range width {
		fmt.Scan(&width[i])
	}

	cases := make([][]int, c)
	for i := range cases {
		cases[i] = make([]int, 2)
		fmt.Scanf("%d %d\n", &cases[i][0], &cases[i][1])
	}

	for _, v := range serviceLane(width, cases) {
		fmt.Println(v)
	}
}
