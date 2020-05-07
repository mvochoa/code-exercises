package main

import (
	"fmt"
)

func beautifulTriplets(d int, a []int) int {
	cont := 0
	for i := range a {
		for j := i+1; j < len(a); j++ {
			if (a[j] - a[i]) == d {
				for k := j+1; k < len(a); k++ {
					if (a[k] - a[j]) == d {
						cont++
					}
				}
			}
		} 
	}
	return cont
}

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(beautifulTriplets(d, a))
}