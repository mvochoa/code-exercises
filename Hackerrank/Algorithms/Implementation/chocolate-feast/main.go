package main

import "fmt"

func chocolateFeast(n,c,m int) int {
	count := int(n / c)
	chocolates := count
	for count >= m {
		count -= m
		count++
		chocolates++
	}
	return chocolates
}

func main() {
	var cases, n,c,m int
	fmt.Scanln(&cases)

	for cases > 0 {
		fmt.Scanf("%d %d %d", &n, &c, &m)
		fmt.Println(chocolateFeast(n,c,m))
		cases--
	}
}
