package main

import (
	"fmt"
)

func candy(ratings []int) int {
	total := 1
	min_rate := 20001 // larger than any possible constraint
	lowest_index := -1

	for i, rate := range ratings {
		if rate < min_rate {
			min_rate = rate
			lowest_index = i
		}
	}

	prev := 1
	for i := lowest_index + 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			prev += 1
		} else {
			prev -= 1
		}

		total += prev
	}

	prev = 1
	for i := lowest_index - 1; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			prev += 1
		} else {
			prev -= 1
		}

		total += prev
	}

	return total
}

func main() {
	for i := 0; i < 0; i++ {
		fmt.Println("balls")
	}
	ratings := []int{1, 3, 2, 2, 1}
	fmt.Println(candy(ratings))
}
