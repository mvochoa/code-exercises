package main

import (
	"bufio"
	"fmt"
	"os"
)

func hello(text string) string {
	return "Hello, World.\n" + text
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print(hello(scanner.Text()))
	}
}
