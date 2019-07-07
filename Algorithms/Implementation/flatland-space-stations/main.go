package main

import (
	"fmt"
	"math"
	"sort"
)

func flatlandSpaceStations(n int, c []int) int {
	var dif, j, max, value int

	sort.Ints(c)
	for i := 0; i < len(c); i++ {
		if i < len(c)-1 {
			dif = int((c[i+1]-c[i])/2) + c[i] + 1
		} else {
			dif = n
		}
		for ; j < dif; j++ {
			value = int(math.Abs(float64(j - c[i])))
			if value > max {
				max = value
			}
		}
	}

	return max
}

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	if n != m {
		c := make([]int, m)
		for i := range c {
			fmt.Scan(&c[i])
		}

		fmt.Println(flatlandSpaceStations(n, c))
	} else {
		fmt.Println(0)
	}
}
