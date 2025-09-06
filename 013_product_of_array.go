package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	front_prod := 1
	back_prod := 1

	for i := 0; i < len(nums); i++ {
		result[i] = 1
		result[i] *= front_prod
		front_prod *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= back_prod
		back_prod *= nums[i]
	}

	return result
}

func main() {
	test := []int{1, 2, 3, 4}

	fmt.Println(productExceptSelf(test))
}
