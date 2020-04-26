package main

import (
	"fmt"
)

func gridSearch(G []string, P []string) string {
	band := false
	var jP, jG int
	lenG, lenP := len(G), len(P)
	leng, lenp := len(G[0]), len(P[0])
	for iG, g := range G {
		for i := 0; i <= leng-lenp; i++ {
			if g[i:i+lenp] == P[0] {
				band = true
				for jP, jG = 1, iG+1; jP < lenP && jG < lenG; jP, jG = jP+1, jG+1 {
					if G[jG][i:i+lenp] != P[jP] {
						band = false
						break
					}
				}
				if band && jG <= lenG {
					return "YES"
				}
			}
		}
	}

	return "NO"
}

func main() {
	var cases, R, r, c int
	fmt.Scanln(&cases)

	for ; cases > 0; cases-- {
		fmt.Scanf("%d %d", &R, &c)
		G := make([]string, R)
		for i := range G {
			fmt.Scan(&G[i])
		}
		fmt.Scanf("%d %d", &r, &c)
		P := make([]string, r)
		for i := range P {
			fmt.Scan(&P[i])
		}
		fmt.Println(gridSearch(G, P))
	}

}
