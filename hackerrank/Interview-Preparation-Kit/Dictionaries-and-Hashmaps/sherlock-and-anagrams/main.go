package main

import (
	"fmt"
)

func contains(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := range s1 {
		m1[s1[i]]++
		m2[s2[i]]++
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

func sherlockAndAnagrams(w string) int {
	m := make(map[string]int)
	s, c, l := 1, 0, len(w)
	for s < l {
		for i := 0; i < l-(s-1); i++ {
			m[w[i:i+s]]++
		}
		s++
	}

	return c
}

func main() {
	var c int
	fmt.Scanln(&c)

	a := make([]string, c)
	for i := range a {
		fmt.Scanf("%s", &a[i])
	}

	for _, v := range a {
		fmt.Println(sherlockAndAnagrams(v))
	}
}
