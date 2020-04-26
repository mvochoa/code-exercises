package main

import (
	"fmt"
)

func plusMinus(size int, arreglo []int) {
	res := make([]int, 3)
	for _, v := range arreglo {
		if v == 0 {
			res[2]++
		} else if v < 0 {
			res[1]++
		} else {
			res[0]++
		}
	}

	for _, v := range res {
		fmt.Printf("%8f\n", float64(v)/float64(size))
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	plusMinus(n, arr)
}
