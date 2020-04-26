package main

import "fmt"

func gradingStudents(grades []int) []int {
	var ten int

	for i, v := range grades {
		if v >= 38 && v < 100 {
			ten = int(v/10) * 10
			if (v % 10) <= 5 {
				ten += 5
			} else {
				ten += 10
			}

			if (ten - v) < 3 {
				grades[i] = ten
			}
		}
	}

	return grades
}

func main() {
	var c int
	fmt.Scanln(&c)

	grades := make([]int, c)
	for i := range grades {
		fmt.Scanln(&grades[i])
	}

	for _, v := range gradingStudents(grades) {
		fmt.Println(v)
	}
}
