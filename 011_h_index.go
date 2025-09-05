package main

import (
	"fmt"
	"sort"
)

func ascending_order(i int, j int) bool {
	return i < j
}

func hIndex(citations []int) int {
	sort.Ints(citations)

	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}

	return 0
}

func main() {
	citations := []int{1, 3, 1}

	fmt.Println(hIndex(citations))
}
