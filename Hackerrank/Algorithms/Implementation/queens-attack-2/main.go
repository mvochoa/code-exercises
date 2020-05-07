package main

import "fmt"

func queensAttack(n int, k int, rowQ int, colQ int, obstacles [][]int) int {
	var ok bool
	total := 0
	obs := make(map[string]bool)
	for _, v := range obstacles {
		obs[fmt.Sprintf("%d %d", v[0], v[1])] = true
	}

	for i := rowQ - 1; i >= 1; i-- {
		if ok = obs[fmt.Sprintf("%d %d", i, colQ)]; !ok {
			total++
		} else {
			i = 0
		}
	}
	for i := rowQ + 1; i <= n; i++ {
		if ok = obs[fmt.Sprintf("%d %d", i, colQ)]; !ok {
			total++
		} else {
			i = n
		}
	}
	for i := colQ - 1; i >= 1; i-- {
		if ok = obs[fmt.Sprintf("%d %d", rowQ, i)]; !ok {
			total++
		} else {
			i = 0
		}
	}
	for i := colQ + 1; i <= n; i++ {
		if ok = obs[fmt.Sprintf("%d %d", rowQ, i)]; !ok {
			total++
		} else {
			i = n
		}
	}

	for i, j := rowQ-1, colQ-1; i >= 1 && j >= 1; i, j = i-1, j-1 {
		if ok = obs[fmt.Sprintf("%d %d", i, j)]; !ok {
			total++
		} else {
			i = 0
			j = 0
		}
	}
	for i, j := rowQ+1, colQ+1; i <= n && j <= n; i, j = i+1, j+1 {
		if ok = obs[fmt.Sprintf("%d %d", i, j)]; !ok {
			total++
		} else {
			i = n
			j = n
		}
	}
	for i, j := rowQ+1, colQ-1; i <= n && j >= 1; i, j = i+1, j-1 {
		if ok = obs[fmt.Sprintf("%d %d", i, j)]; !ok {
			total++
		} else {
			i = n
			j = 0
		}
	}
	for i, j := rowQ-1, colQ+1; i >= 1 && j <= n; i, j = i-1, j+1 {
		if ok = obs[fmt.Sprintf("%d %d", i, j)]; !ok {
			total++
		} else {
			i = 0
			j = n
		}
	}

	return total
}

func main() {
	var n, k, r, c int
	fmt.Scanf("%d %d", &n, &k)
	fmt.Scanf("%d %d", &r, &c)

	o := make([][]int, k)
	for i := range o {
		o[i] = make([]int, 2)
		fmt.Scanf("%d %d", &o[i][0], &o[i][1])
	}

	fmt.Println(queensAttack(n, k, r, c, o))
}
