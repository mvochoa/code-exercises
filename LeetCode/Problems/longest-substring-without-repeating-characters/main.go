package main

func lengthOfLongestSubstring(s string) int {
	var ls, run int
	letters := map[rune]int{}

	for i, c := range s {
		if old, ok := letters[c]; ok {
			run = max(old, run)
		}
		ls = max(ls, i+1-run)
		letters[c] = i + 1
	}

	return ls
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
