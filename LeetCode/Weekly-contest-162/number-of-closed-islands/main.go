package main

func cancelLand(grid [][]int, r, c, v int) [][]int {
	if r-1 >= 0 {
		if grid[r-1][c] == 0 {
			grid[r][c] = v
			grid[r-1][c] = v
			grid = cancelLand(grid, r-1, c, v)
		}
	}
	if r+1 <= len(grid)-1 {
		if grid[r+1][c] == 0 {
			grid[r][c] = v
			grid[r+1][c] = v
			grid = cancelLand(grid, r+1, c, v)
		}
	}
	if c-1 >= 0 {
		if grid[r][c-1] == 0 {
			grid[r][c] = v
			grid[r][c-1] = v
			grid = cancelLand(grid, r, c-1, v)
		}
	}
	if c+1 <= len(grid[0])-1 {
		if grid[r][c+1] == 0 {
			grid[r][c] = v
			grid[r][c+1] = v
			grid = cancelLand(grid, r, c+1, v)
		}
	}

	return grid
}

func closedIsland(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	var land int
	count := map[int]int{}
	a := []int{0, len(grid) - 1}
	b := []int{0, len(grid[0]) - 1}

	for row := 1; row < a[1]; row++ {
		for col := 1; col < b[1]; col++ {
			if grid[row][col] == 0 {
				if grid[row-1][col] == 1 && grid[row+1][col] == 1 && grid[row][col-1] == 1 && grid[row][col+1] == 1 {
					grid[row][col] = -1
					land++
				} else {
					grid = cancelLand(grid, row, col, len(count)+2)
					count[len(count)+2] = 1
				}
			}
		}
	}

	var ok bool
	for i := 0; i < len(a) && len(count) > 0; i++ {
		for col := 1; col < b[1]; col++ {
			if _, ok = count[grid[a[i]][col]]; ok {
				delete(count, grid[a[i]][col])
			}
		}
	}

	for i := 0; i < len(b) && len(count) > 0; i++ {
		for row := 1; row < a[1]; row++ {
			if _, ok = count[grid[row][b[i]]]; ok {
				delete(count, grid[row][b[i]])
			}
		}
	}

	return land + len(count)
}

func main() {}
