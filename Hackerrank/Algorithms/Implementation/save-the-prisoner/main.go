package main

import (
	"fmt"
)

func saveThePrisoner(n int, m int, s int) int {
	v := (m % n)
	s = v + (s - 1)
	if s == 0 {
		return n
	} else if s > n {
		return s % n
	}
	return s
}

func main() {
	var i, n, m, s int
	fmt.Scanln(&i)

	for ; i > 0; i-- {
		fmt.Scanf("%d %d %d\n", &n, &m, &s)
		fmt.Println(saveThePrisoner(n, m, s))
	}
}
