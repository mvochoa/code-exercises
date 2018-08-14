package main

import (
	"fmt"
	"time"
)

func countTriplets(a []int64, r int64) int64 {
	var n1, n2, c int64 = 0, 0, 0
	var j, k []int
	var w, p int
	var ok bool
	m := make(map[int64][]int)

	for i, v := range a {
		m[v] = append(m[v], i)
	}

	for i := 0; i < len(a); i++ {
		n1 = a[i] * r
		n2 = n1 * r
		if j, ok = m[n1]; ok {
			k, ok = m[n2]
			if ok {
				w = func(q int) int {
					for l := q; l < len(j); l++ {
						if j[l] > i {
							return l
						}
					}
					return len(j)
				}(0)
				p = 0
				for ; w < len(j); w++ {
					p = func(q int) int {
						for l := q; l < len(k); l++ {
							if k[l] > j[w] {
								return l
							}
						}
						return len(k)
					}(p)
					c += int64(len(k) - p)
				}
			}
		}
	}

	return c
}

func main() {
	start := time.Now()
	var n int
	var r int64
	fmt.Scanf("%d %d\n", &n, &r)

	a := make([]int64, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(countTriplets(a, r))
	e := time.Now().Sub(start)
	fmt.Println(e)
}
