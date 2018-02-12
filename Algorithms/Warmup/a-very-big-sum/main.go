package main

import "fmt"

func aVeryBigSum(x []uint64) {
	var sum uint64
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	var a int
	fmt.Scanf("%v", &a)

	x := make([]uint64, a)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}

	n, _ := fmt.Scanln(y...)
	x = x[:n]

	aVeryBigSum(x)
}
