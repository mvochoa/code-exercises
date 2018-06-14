package main

import (
	"fmt"
	"math"
)

var matriz = [][]int{
	[]int{2, 7, 6},
	[]int{9, 5, 1},
	[]int{4, 3, 8},
}

func formingMagicSquare(m [3][3]int) int {
	var cost [8]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			addCost(&cost[0], m[i][j], matriz[i][j])
			addCost(&cost[1], m[i][j], matriz[i][2-j])
			addCost(&cost[2], m[i][j], matriz[2-i][j])
			addCost(&cost[3], m[i][j], matriz[2-i][2-j])
			addCost(&cost[4], m[i][j], matriz[j][i])
			addCost(&cost[5], m[i][j], matriz[j][2-i])
			addCost(&cost[6], m[i][j], matriz[2-j][i])
			addCost(&cost[7], m[i][j], matriz[2-j][2-i])
		}
	}

	for _, v := range cost {
		if v < cost[0] {
			cost[0] = v
		}
	}

	return cost[0]
}

func addCost(cost *int, a, b int) {
	*cost += int(math.Abs(float64(a - b)))
}

func main() {
	var m [3][3]int
	fmt.Scanf("%d %d %d\n%d %d %d\n%d %d %d\n",
		&m[0][0],
		&m[0][1],
		&m[0][2],
		&m[1][0],
		&m[1][1],
		&m[1][2],
		&m[2][0],
		&m[2][1],
		&m[2][2])

	fmt.Println(formingMagicSquare(m))
}
