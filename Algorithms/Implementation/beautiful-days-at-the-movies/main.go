package main

import (
	"fmt"
	"math"
	"strconv"
)

func beautifulDays(i, j, k int) int {
	var s string
	var n float64
	days := 0
	for a := i; a <= j; a++ {
		s = ""
		for i = a; i > 0; {
			s += strconv.Itoa(i % 10)
			i = int(i / 10)
		}

		n, _ = strconv.ParseFloat(s, 64)
		n := math.Abs(float64(a)-n) / float64(k)
		if float64(int(n)) == n {
			days++
		}
	}

	return days
}

func main() {
	var i, j, k int
	fmt.Scanf("%d %d %d\n", &i, &j, &k)

	fmt.Println(beautifulDays(i, j, k))
}
