package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	base := strs[0]
	prefix_len := len(base)

	for _, target := range strs {
		if prefix_len > len(target) {
			prefix_len = len(target)
		}

		for i := 0; i < prefix_len; i++ {
			if base[i] != target[i] {
				prefix_len = i
				break
			}
		}
	}

	return base[0:prefix_len]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
