package main

import (
	"fmt"
)

func jumpingOnClouds(c []int, k int) int {
	e := 100
	for i := 0; i < len(c); i, e = i+k, e-1 {
		if c[i] == 1 {
			e -= 2
		}
	}

	return e
}

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	clouds := make([]int, n)
	for i := range clouds {
		fmt.Scan(&clouds[i])
	}

	fmt.Println(jumpingOnClouds(clouds, k))
}
