package main

import (
	"fmt"
)

func jump(nums []int) int {
	jumps := 0
	last_jump_index := 0
	max_index := 0

	for i := 0; i < len(nums); i++ {
		if i > last_jump_index {
			jumps++
			last_jump_index = max_index
		}

		if nums[i]+i > max_index {
			max_index = nums[i] + i
		}
	}

	return jumps
}

func main() {
	input := []int{2, 3, 0, 1, 4}

	fmt.Println(jump(input))
}
