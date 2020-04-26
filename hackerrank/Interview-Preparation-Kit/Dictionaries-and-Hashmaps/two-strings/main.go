package main

import (
	"fmt"
)

func twoStrings(s1 string, s2 string) string {
	m := make(map[byte]bool)
	for i := range s1 {
		m[s1[i]] = false
	}

	for i := range s2 {
		if _, ok := m[s2[i]]; ok {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	var c int
	fmt.Scanln(&c)

	s := make([][]string, c)
	for i := range s {
		s[i] = make([]string, 2)
		for j := range s[i] {
			fmt.Scanf("%s", &s[i][j])
		}
	}

	for _, v := range s {
		fmt.Println(twoStrings(v[0], v[1]))
	}
}
