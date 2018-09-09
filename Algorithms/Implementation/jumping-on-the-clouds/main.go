package main

import "fmt"

func jumpingOnClouds(c []int) int {
	j, i := 0, 0
	for ; i < len(c)-2; i++ {
		if c[i+2] == 0 {
			i++
		}
		j++
	}
	if i < len(c)-1 && c[len(c)-1] != 1 {
		j++
	}

	return j
}

func main() {
	var n int
	fmt.Scanln(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(jumpingOnClouds(a))
}
