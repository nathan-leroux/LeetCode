package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	var candidate int
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	mymap := []int{1, 2, 2, 1, 2}

	fmt.Println(majorityElement(mymap))
	fmt.Println("yeehaw")
}
