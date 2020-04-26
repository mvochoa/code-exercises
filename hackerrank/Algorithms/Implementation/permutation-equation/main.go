package main

import (
	"fmt"
)

func permutationEquation(p []int) {
	a := make([]int, len(p)+1)
	for i, v := range p {
		a[v] = i + 1
	}

	for i := 1; i <= len(p); i++ {
		fmt.Println(a[a[i]])
	}
}

func main() {
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	permutationEquation(arr)
}
