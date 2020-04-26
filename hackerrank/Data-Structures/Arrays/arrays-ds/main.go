package main

import "fmt"

func arraysDS(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		fmt.Printf("%d ", a[i])
	}
	fmt.Print(a[0])
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	arraysDS(a)
}
