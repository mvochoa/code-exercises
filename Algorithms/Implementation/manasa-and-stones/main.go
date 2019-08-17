package main

import (
	"fmt"
	"sort"
	"strings"
)

func stones(n int, a int, b int) []int {
	stones := make(map[int]bool)
	for i := 0; i < n; i++ {
		stones[a*((n-1)-i)+b*i] = true
	}

	var result []int
	for k := range stones {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

func main() {
	var cases, n, a, b int
	fmt.Scanln(&cases)

	var values []string
	for cases > 0 {
		fmt.Scanln(&n)
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		values = append(values, strings.Trim(fmt.Sprint(stones(n, a, b)), "[|]"))
		cases--
	}

	for _, v := range values {
		fmt.Println(v)
	}
}
