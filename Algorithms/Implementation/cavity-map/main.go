package main

import (
	"fmt"
	"strconv"
)

func cavityMap(grid []string) []string {
	m := make([][]int, len(grid))
	for i := range m {
		m[i] = make([]int, len(grid[i]))
		for j := range grid[i] {
			m[i][j], _ = strconv.Atoi(string(grid[i][j]))
		}
	}

	for row := 1; row < len(m)-1; row++ {
		for col := 1; col < len(m[row])-1; col++ {
			n := m[row][col]
			if n > m[row-1][col] && n > m[row+1][col] && n > m[row][col-1] && n > m[row][col+1] {
				grid[row] = grid[row][:col] + "X" + grid[row][col+1:]
			}
		}
	}

	return grid
}

func main() {
	var n int
	fmt.Scanln(&n)

	grid := make([]string, n)
	for i := range grid {
		fmt.Scanln(&grid[i])
	}

	for _, v := range cavityMap(grid) {
		fmt.Println(v)
	}
}
