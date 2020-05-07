package main

import (
	"fmt"
)

func libraryFine(d, m, y, d2, m2, y2 int) int {
	y -= y2
	m -= m2
	d -= d2

	if y < 0 || (y <= 0 && m < 0) {
		return 0
	}

	if y > 0 {
		return 10000
	} else if m > 0 {
		return 500 * m
	} else if d > 0 && m >= 0 {
		return 15 * d
	}
	return 0
}

func main() {
	var d, m, y int
	var d1, m1, y1 int
	fmt.Scanf("%d %d %d\n", &d, &m, &y)
	fmt.Scanf("%d %d %d\n", &d1, &m1, &y1)

	fmt.Println(libraryFine(d, m, y, d1, m1, y1))
}
