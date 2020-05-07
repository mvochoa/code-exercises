package main

import (
	"fmt"
	"sort"
)

func acmTeam(topic [][]int) [2]int {
	var r [2]int
	var calc []int
	var count int

	for d := 0; d < len(topic)-1; d++ {
		for i := d + 1; i < len(topic); i++ {
			count = 0
			for j := 0; j < len(topic[i]); j++ {
				if topic[d][j] == 1 || topic[i][j] == 1 {
					count++
				}
			}
			calc = append(calc, count)
		}
	}

	sort.Ints(calc)
	if len(calc) > 0 {
		r[0] = calc[len(calc)-1]
		r[1]++
		for i := len(calc) - 2; i >= 0; i-- {
			if r[0] == calc[i] {
				r[1]++
			}
		}
	}

	return r
}

func main() {
	var s string
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	t := make([][]int, n)
	for i := range t {
		t[i] = make([]int, m)
		fmt.Scan(&s)
		for j, v := range s {
			if '1' == v {
				t[i][j] = 1
			}
		}
	}

	for _, v := range acmTeam(t) {
		fmt.Println(v)
	}
}
