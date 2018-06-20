package main

import (
	"fmt"
)

func climbingLeaderboard(scores []int, alice []int) []int {
	sizeS, sizeA, p := len(scores), len(alice), 1
	r := make([]int, sizeA)

	ant := scores[0]
	for i, k := sizeA-1, 0; i >= 0; i-- {
		for ; k < sizeS; k++ {
			if scores[k] != ant {
				p++
			}
			ant = scores[k]

			if alice[i] >= scores[k] {
				r[i] = p
				break
			}

		}
	}

	for i := 0; i < sizeA && r[i] == 0; i++ {
		r[i] = p + 1
	}

	return r
}

func main() {
	var n, m int
	fmt.Scanln(&n)
	scores := scanArray(n)

	fmt.Scanln(&m)
	alice := scanArray(m)

	for _, v := range climbingLeaderboard(scores, alice) {
		fmt.Println(v)
	}
}

func scanArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	return arr
}
