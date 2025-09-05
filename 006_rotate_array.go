package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	// create second array
	// write with index offset
	nums_copy := make([]int, len(nums))
	copy(nums_copy, nums)

	var index int

	for i := 0; i < len(nums); i++ {
		index = i + k
		if index >= len(nums) {
			index %= len(nums)
		}

		nums[index] = nums_copy[i]
	}
}
