package main

import (
	"fmt"
	"math"
)

func diagonalDifference(size int, arreglo [][]int) {
	var sumD, sumI int
	for i := 0; i < size; i++ {
		sumD += arreglo[i][i]
		sumI += arreglo[i][size-i-1]
	}
	fmt.Println(math.Abs(float64(sumD - sumI)))
}

func capt(s int) []int {
	x := make([]int, s)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, _ := fmt.Scanln(y...)
	x = x[:n]
	return x
}

func main() {
	var a int
	fmt.Scanf("%v", &a)

	b := make([][]int, a)
	for i := range b {
		b[i] = capt(a)
	}

	diagonalDifference(a, b)
}
