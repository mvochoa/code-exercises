package main

import (
	"fmt"
	"sort"
)

func organizingContainers(container [][]int) string {
	msg := "Possible"
	array := make([][]int, 2)
	array[0] = make([]int, len(container))
	array[1] = make([]int, len(container))

	for i, v := range container {
		for j, u := range v {
			array[0][i] += u
			array[1][j] += u
		}
	}

	sort.Ints(array[0])
	sort.Ints(array[1])
	for i := range array[0] {
		if array[0][i] != array[1][i] {
			msg = "Impossible"
			break
		}
	}

	return msg
}

func main() {
	var numCases, size, num int

	fmt.Scanln(&numCases)
	cases := make([][][]int, numCases)
	for c := range cases {
		fmt.Scanln(&size)
		cases[c] = make([][]int, size)
		for i := range cases[c] {
			cases[c][i] = make([]int, size)

			for j := range cases[c][i] {
				fmt.Scan(&num)
				cases[c][i][j] = num
			}
		}

	}

	for _, v := range cases {
		fmt.Println(organizingContainers(v))
	}
}
