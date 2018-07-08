package main

import (
	"fmt"
)

func findDigits(n int) int {
	count := 0
	number := n
	digit := make([]int, 10)
	for n > 0 {
		digit[(n%10)]++
		n = int(n / 10)
	}

	for i := 1; i < 10; i++ {
		if digit[i] > 0 && number%i == 0 {
			count += digit[i]
		}
	}

	return count
}

func main() {
	var c, n int
	fmt.Scan(&c)

	for c > 0 {
		fmt.Scanln(&n)
		fmt.Println(findDigits(n))
		c--
	}
}
