package main

import (
	"fmt"
)

func swap(val1 *int, val2 *int) {
	int_copy := *val1
	*val1 = *val2
	*val2 = int_copy
}

func removeElement(nums []int, val int) int {
	total := len(nums)
	tail := 0

	for i := 0; i < total; i++ {
		if nums[i] != val {
			swap(&nums[i], &nums[tail])
			tail++
		}
	}

	return tail
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)

	fmt.Println(k)
	fmt.Println(nums)
}
