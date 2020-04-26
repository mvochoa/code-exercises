package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	iP, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	dP, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	s += scanner.Text()

	fmt.Println(iP + i)
	fmt.Printf("%.1f\n", dP+float64(uint64(d)))
	fmt.Print(s)

}
