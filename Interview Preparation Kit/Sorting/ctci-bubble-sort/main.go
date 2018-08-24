package main

import "fmt"

func countSwaps(a []int) {
	var count int
	var band bool
	l, aux := len(a)-1, 0
	for range a {
		band = true
		for i := 0; i < l; i++ {
			if a[i] > a[i+1] {
				aux = a[i]
				a[i] = a[i+1]
				a[i+1] = aux
				band = false
				count++
			}
		}

		if band {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", count)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d", a[l])
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	countSwaps(a)
}
