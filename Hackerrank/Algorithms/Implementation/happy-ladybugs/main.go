package main

import (
	"fmt"
)

func happyLadybugs(b string) string {
	m := make(map[string]int)
	for _, v := range b {
		m[string(v)]++
	}

	for k, v := range m {
		if v == 1 && k != "_" {
			return "NO"
		}
	}

	if _, ok := m["_"]; !ok {
		for i := len(b) - 2; i > 0; i-- {
			if b[i] != b[i+1] && b[i] != b[i-1] {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	var n int
	var b string
	fmt.Scanln(&n)

	for i := n - 1; i >= 0; i-- {
		fmt.Scanln(&n)
		fmt.Scanln(&b)
		fmt.Println(happyLadybugs(b))
	}
}
