package main

import (
	"fmt"
)

func checkMagazine(magazine, note []string) {
	var band bool
	size := len(magazine)-1
	for _, v := range note {
		band = false
		for i := size; i >= 0; i--{
			if v == magazine[i] {
				magazine[i] = ""
				band = true
				i = 0
			}
		}
		if !band {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	a := make([]string, n)
	b := make([]string, m)
	for i := range a {
		fmt.Scanf("%s", &a[i])
	}
	for i := range b {
		fmt.Scanf("%s", &b[i])
	}

	checkMagazine(a, b)
}
