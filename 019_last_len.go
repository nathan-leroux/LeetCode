package main

import (
	"fmt"
	"unicode"
)

func lengthOfLastWord(s string) int {
	var end int
	start := -1

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			end = i
			break
		}
	}

	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			start = i
			break
		}
	}

	return end - start
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
