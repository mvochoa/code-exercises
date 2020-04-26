package main

import (
	"fmt"
	"sort"
)

func hurdleRace(k int, height []int) int {
	max := 0
	sort.Ints(height)
	if len(height) > 0 {
		if height[len(height)-1] > k {
			max = (height[len(height)-1] - k)
		}
	}

	return max
}

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	heights := make([]int, n)
	for i := range heights {
		fmt.Scan(&heights[i])
	}

	fmt.Println(hurdleRace(k, heights))
}
