package main

import "fmt"

func taumBday(b int, w int, bc int, wc int, z int) int {
	res := (b * bc) + (w * wc)

	if o := ((b * bc) + (w * (bc + z))); o < res {
		res = o
	}
	if o := ((w * wc) + (b * (wc + z))); o < res {
		res = o
	}

	return res
}

func main() {
	var t int
	fmt.Scanln(&t)

	cases := make([][]int, t)
	for i := range cases {
		cases[i] = make([]int, 5)
		fmt.Scanf("%d %d\n", &cases[i][0], &cases[i][1])
		fmt.Scanf("%d %d %d\n", &cases[i][2], &cases[i][3], &cases[i][4])
	}

	for _, v := range cases {
		fmt.Println(taumBday(v[0], v[1], v[2], v[3], v[4]))
	}
}
