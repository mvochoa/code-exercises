package main

import (
	"fmt"
)

func viralAdvertising(n int) int {
	cumulative, s, l := 0, 6, 2
	for day := n; day > 0; day-- {
		cumulative += l
		l = int(s / 2)
		s = l * 3
	}

	return cumulative
}

func main() {
	var n int
	fmt.Scanln(&n)

	fmt.Println(viralAdvertising(n))
}
