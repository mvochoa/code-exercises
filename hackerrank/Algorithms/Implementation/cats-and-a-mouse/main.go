package main

import (
	"fmt"
	"math"
)

func catAndMouse(x, y, z int) string {
	g1 := math.Abs(float64(x - z))
	g2 := math.Abs(float64(y - z))
	if g1 == g2 {
		return "Mouse C"
	} else if g1 < g2 {
		return "Cat A"
	} else {
		return "Cat B"
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	array := make([][3]int, n)
	for i := 0; i < n; i++ {
		var numbers [3]int
		fmt.Scanf("%d %d %d\n", &numbers[0], &numbers[1], &numbers[2])
		array[i] = numbers
	}

	for _, item := range array {
		fmt.Println(catAndMouse(item[0], item[1], item[2]))
	}
}
