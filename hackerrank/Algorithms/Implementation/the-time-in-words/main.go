package main

import "fmt"

var numbers = map[int]string{
	0:  "",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
}

func getMinutes(m int) string {
	var minutes string
	if m > 20 {
		minutes = numbers[(int(m/10) * 10)]
		if u := m % 10; u > 0 {
			minutes += " " + numbers[u]
		}
	} else {
		minutes = numbers[m]
	}

	if minutes == "one" {
		return fmt.Sprintf("%s minute", minutes)
	}

	return fmt.Sprintf("%s minutes", minutes)
}

func timeInWords(h, m int) string {
	if m == 0 {
		return fmt.Sprintf("%s o' clock", numbers[h])
	} else if m == 15 {
		return fmt.Sprintf("quarter past %s", numbers[h])
	} else if m == 30 {
		return fmt.Sprintf("half past %s", numbers[h])
	} else if m == 45 {
		return fmt.Sprintf("quarter to %s", numbers[h+1])
	}

	if m < 30 {
		return fmt.Sprintf("%s past %s", getMinutes(m), numbers[h])
	}
	return fmt.Sprintf("%s to %s", getMinutes(60-m), numbers[h+1])
}

func main() {
	var h, m int
	fmt.Scanln(&h)
	fmt.Scanln(&m)

	fmt.Println(timeInWords(h, m))
}
