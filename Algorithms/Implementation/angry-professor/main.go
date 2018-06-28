package main

import (
	"fmt"
)

func angryProfessor(k int, a []int) string {
	s := 0
	for _, v := range a {
		if v <= 0 {
			s++
		}
	}

	if s >= k {
		return "NO"
	}

	return "YES"
}

func main() {
	var n, t, k int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &t, &k)
		arr := make([]int, t)
		for i := range arr {
			fmt.Scan(&arr[i])
		}
		fmt.Println(angryProfessor(k, arr))
	}
}
