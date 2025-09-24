package main

import (
	"fmt"
)

func convertRoman(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		fmt.Println("unexpected input: ", c)
		return -1
	}
}

func romanToInt(s string) int {
	total := 0
	var value int
	var next int

	for i := 0; i < len(s); i++ {
		value = convertRoman(rune(s[i]))

		if i >= len(s)-1 {
			total += value
			continue
		}

		next = convertRoman(rune(s[i+1]))
		if next > value {
			total += next - value
			i++
			continue
		}

		total += value
	}
	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
