package main

import (
	"fmt"
)

func swap(val1 *int, val2 *int) {
	int_copy := *val1
	*val1 = *val2
	*val2 = int_copy
}

func removeDuplicates(nums []int) int {
	count := 1
	previous := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != previous {
			previous = nums[i]
			swap(&nums[count], &nums[i])
			count++
		}
	}

	return count
}

func main() {
	nums := []int{0, 0, 1, 1, 2}

	unique := removeDuplicates(nums)

	fmt.Printf("unique: %d\nnums: ", unique)
	fmt.Println(nums)

}
