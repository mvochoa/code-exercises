package main

import (
	"fmt"
	"math"
	"strings"
)

func encryption(s string) string {
	// Calculate matrix size
	size := math.Sqrt(float64(len(s)))
	rows := int(size)
	columns := int(size)
	if (size - float64(int(size))) > 0 {
		columns++
	}
	if (rows * columns) < len(s) {
		rows++
	}

	// Put the string in the matrix.
	position := 0
	matrix := make([][]string, rows)
	for i := range matrix {
		if (position + columns) < len(s) {
			matrix[i] = strings.Split(s[position:position+columns], "")
		} else {
			matrix[i] = make([]string, columns)
			for j, d := position, 0; j < len(s); j, d = j+1, d+1 {
				matrix[i][d] = string(s[j])
			}
		}
		position += columns
	}

	// Rebuild the string using the matrix.
	var rebuiltString string
	for j := 0; j < columns; j++ {
		for i := 0; i < rows; i++ {
			rebuiltString += matrix[i][j]
		}
		rebuiltString += " "
	}

	return rebuiltString
}

func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Println(encryption(s))
}
