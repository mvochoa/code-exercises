package main

import "fmt"

func howManyGames(p, d, m, s int) {
	var count int
	for (s - p) >= 0 {
		s -= p
		if (p - d) > m {
			p -= d
		} else {
			p = m
		}
		count++
	}

	fmt.Println(count)
}

func main() {
	var p, d, m, s int
	fmt.Scanf("%d %d %d %d", &p, &d, &m, &s)

	howManyGames(p, d, m, s)
}
