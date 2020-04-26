package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int64) {
	f := big.NewInt(1)
	for n > 0 {
		f = f.Mul(f, big.NewInt(n))
		n--
	}
	fmt.Println(f.String())
}

func main() {
	var n int64
	fmt.Scanln(&n)

	extraLongFactorials(n)
}
