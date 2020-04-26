package main

import (
	"fmt"
)

func designerPdfViewer(h map[string]int, word string) int {
	height := 0
	for _, w := range word {
		if h[string(w)] > height {
			height = h[string(w)]
		}
	}

	return len(word) * height
}

func main() {
	var word string
	var n int
	chars := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, v := range keys {
		fmt.Scan(&n)
		chars[v] = n
	}
	fmt.Scanln(&word)

	fmt.Println(designerPdfViewer(chars, word))
}
