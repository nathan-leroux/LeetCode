package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	delta := 0
	lowest := 10001
	lowest_index := -1

	for i := 0; i < len(gas); i++ {
		delta += gas[i] - cost[i]
		if delta < lowest {
			lowest = delta
			lowest_index = i
		}
	}

	if delta < 0 {
		return -1
	}

	if lowest_index == len(gas)-1 {
		return 0
	}

	return lowest_index + 1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
