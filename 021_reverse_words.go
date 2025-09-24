package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	slices.Reverse(words)

	result := words[0]

	for i := 1; i < len(words); i++ {
		// fmt.Printf("i:%d, string: '%s'\n", i, words[i])
		if words[i] == "" {
			// fmt.Printf("skipped i:%d\n", i)
			continue
		}

		result = result + " " + words[i]
	}

	return strings.Trim(result, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
