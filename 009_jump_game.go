package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	max_index := 0

	for i := 0; i < len(nums); i++ {
		if i > max_index {
			return false
		}

		if nums[i]+i > max_index {
			max_index = nums[i] + i
		}
	}

	return true
}

func main() {
}
