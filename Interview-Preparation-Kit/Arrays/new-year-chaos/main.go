package main

import (
	"fmt"
)

func minimumBribes(q []int) {
	s, b, aux := 0, 0, 0
	band := true
	for i := 0; i < len(q); i++ {
		b = q[i] - (i + 1)
		if b < 0 {
			for ; b < 0; b, i, s = b+1, i-1, s+1 {
				aux = q[i]
				q[i] = q[i-1]
				q[i-1] = aux
			}
		} else if b > 2 {
			band = false
			i = len(q)
		}
	}

	if band {
		fmt.Println(s)
	} else {
		fmt.Println("Too chaotic")
	}
}

func main() {
	var c, n int
	fmt.Scanln(&c)

	cases := make([][]int, c)
	for c = range cases {
		fmt.Scanln(&n)
		cases[c] = make([]int, n)
		for i := range cases[c] {
			fmt.Scan(&cases[c][i])
		}
	}

	for _, v := range cases {
		minimumBribes(v)
	}
}
