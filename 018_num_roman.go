package main

import (
	"fmt"
)

func intToRoman(num int) string {
	var digit int
	var roman_char string
	power := 0
	result := ""

	romans := map[int]map[int]string{
		0: {
			1: "I",
			5: "V",
		},
		1: {
			1: "X",
			5: "L",
		},
		2: {
			1: "C",
			5: "D",
		},
		3: {
			1: "M",
		},
	}

	for num != 0 {
		digit = num % 10
		roman_char = ""

		if digit == 4 {
			roman_char = romans[power][1] + romans[power][5]

		} else if digit == 9 {
			roman_char = romans[power][1] + romans[power+1][1]

		} else {
			if digit/5 >= 1 {
				roman_char = romans[power][5]
				digit -= 5
			}

			for i := 0; i < digit; i++ {
				roman_char += romans[power][1]
			}
		}

		power++
		num = num / 10
		result = roman_char + result
	}

	return result
}

func main() {
	fmt.Println(intToRoman(1997))
}
