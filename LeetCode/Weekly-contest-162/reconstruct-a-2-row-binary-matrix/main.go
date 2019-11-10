package main

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var p []int
	a := make([]int, len(colsum))
	b := make([]int, len(colsum))

	for i, v := range colsum {
		if v == 2 && upper > 0 && lower > 0 {
			a[i] = 1
			b[i] = 1
			upper--
			lower--
		} else if v == 1 {
			p = append(p, i)
		} else if v != 0 {
			return [][]int{}
		}
	}

	for _, i := range p {
		if upper > 0 {
			a[i] = 1
			upper--
		} else if lower > 0 {
			b[i] = 1
			lower--
		} else {
			return [][]int{}
		}
	}

	if upper > 0 || lower > 0 {
		return [][]int{}
	}

	return [][]int{a, b}
}

func main() {}
