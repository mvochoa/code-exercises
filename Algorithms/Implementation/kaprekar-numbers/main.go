package main

import (
	"fmt"
	"strconv"
	"strings"
)

func kaprekarNumbers(p, q int) {
	var nums []string
	var n string
	var l, r int

	for i := p; i <= q; i++ {
		n = strconv.Itoa((i * i))
		l, _ = strconv.Atoi(n[:int(len(n)/2)])
		r, _ = strconv.Atoi(n[int(len(n)/2):])
		if (l + r) == i {
			nums = append(nums, strconv.Itoa(i))
		}
	}

	if len(nums) != 0 {
		fmt.Println(strings.Join(nums, " "))
	} else {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	var p, q int
	fmt.Scanln(&p)
	fmt.Scanln(&q)

	kaprekarNumbers(p, q)
}
