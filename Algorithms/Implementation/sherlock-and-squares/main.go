package main

import (
	"fmt"
	"math"
)

func squares(a int, b int) int {
	r, c := 0, 0
	for i := int(math.Sqrt(float64(a))); r <= b; i++ {
		r = i * i
		if r >= a && r <= b {
			c++
		}
	}
	return c
}

func main() {
	var t, a, b int
	fmt.Scanln(&t)

	for t > 0 {
		fmt.Scanf("%d %d\n", &a, &b)
		fmt.Println(squares(a, b))
		t--
	}
}
