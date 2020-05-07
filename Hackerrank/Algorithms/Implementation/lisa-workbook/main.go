package main

import "fmt"

func workbook(n, k int, arr []int) int {
	page := 1
	special := 0
	band := true
	for _, p := range arr {
		for i := 1; i <= p; i++ {
			band = false
			if i == page {
				special++
			}
			if i%k == 0 {
				band = true
				page++
			}
		}
		if !band {
			page++
		}
	}

	return special
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(workbook(n, k, a))
}
