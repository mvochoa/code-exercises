package main

func oddCells(n int, m int, indices [][]int) int {
	var odd, a int
	r := map[int]int{}
	c := map[int]int{}
	for _, v := range indices {
		r[v[0]]++
		c[v[1]]++
	}

	for i := 0; i < n; i++ {
		a = 0
		if r[i]%2 != 0 {
			a = 1
			odd += m
		}
		for _, u := range c {
			if u%2 != 0 && a == 1 {
				odd--
			} else if u%2 != 0 {
				odd++
			}
		}
	}

	return odd
}

func main() {}
