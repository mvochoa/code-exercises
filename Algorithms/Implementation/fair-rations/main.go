package main

import "fmt"

func fairRations(B []int) int {
	var band bool
	breads := 0
	for i := 0; i < len(B); i++ {
		if B[i]%2 != 0 {
			B[i]++
			breads += 2
			band = true
			if i > 0 {
				band = !(B[i-1]%2 != 0)
				if !band {
					B[i-1]++
				}
			}
			if i < len(B)-1 {
				if band {
					B[i+1]++
				}
			} else if band {
				B[i-1]++
			}
		}
	}

	for _, v := range B {
		if v%2 != 0 {
			return -1
		}
	}

	return breads
}

func main() {
	var n int
	fmt.Scanln(&n)

	b := make([]int, n)
	for i := range b {
		fmt.Scan(&b[i])
	}

	if breads := fairRations(b); breads != -1 {
		fmt.Println(breads)
	} else {
		fmt.Println("NO")
	}
}
