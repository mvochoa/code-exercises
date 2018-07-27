package main

import (
	"fmt"
)

func hourglassSum(arr [][]int) int {
	sum, s := -63, 0

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			s = arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]

			if s > sum {
				sum = s
			}
		}
	}

	return sum
}

func main() {
	arr := make([][]int, 6)

	for i := range arr {
		arr[i] = make([]int, 6)
		for j := range arr[i] {
			fmt.Scan(&arr[i][j])
		}
	}

	fmt.Println(hourglassSum(arr))
}
