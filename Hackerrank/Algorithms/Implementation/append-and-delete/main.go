package main

import (
	"fmt"
)

func appendAndDelete(s string, t string, k int) string {
	size := 0

	for i := 0; i < len(s) && i < len(t); i++ {
		if s[i] == t[i] {
			size++
		} else {
			break
		}
	}
	if (len(s) + len(t) - 2*size) > k {
		return "No"
	} else if (len(s)+len(t)-2*size)%2 == k%2 {
		return "Yes"
	} else if (len(s) + len(t) - k) < 0 {
		return "Yes"
	}

	return "No"
}

func main() {
	var k int
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	fmt.Scanln(&k)

	fmt.Println(appendAndDelete(s, t, k))
}
