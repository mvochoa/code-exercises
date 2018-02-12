package main

import "fmt"

func comparetheTriplets(a [3]int, b [3]int) {
	var ta, tb int
	for i := range a {
		if a[i] > b[i] {
			ta++
		} else if b[i] > a[i] {
			tb++
		}
	}
	fmt.Printf("%v %v", ta, tb)
}

func main() {
	var a, b [3]int
	fmt.Scanf("%v %v %v", &a[0], &a[1], &a[2])
	fmt.Scanf("%v %v %v", &b[0], &b[1], &b[2])

	comparetheTriplets(a, b)
}
