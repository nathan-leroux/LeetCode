// 001 merge sorted array

package main

import (
	"fmt"
)

func merge_001(nums1 []int, m int, nums2 []int, n int) {
	nums1_copy := make([]int, m)

	i := 0
	j := 0

	copy(nums1_copy, nums1)
	clear(nums1)
	index := 0

	for i < m || j < n {
		if i == m {
			nums1[index] = nums2[j]
			index++
			j++

		} else if j == n {
			nums1[index] = nums1_copy[i]
			index++
			i++
		} else if nums1_copy[i] < nums2[j] {
			nums1[index] = nums1_copy[i]
			index++
			i++
		} else {
			nums1[index] = nums2[j]
			index++
			j++
		}
	}
}

func main() {
	m := 5
	n := 3
	nums1 := make([]int, m+n)
	nums1 = []int{1, 1, 3, 5, 7}
	nums2 := []int{2, 4, 6}

	merge_001(nums1, m, nums2, n)

	fmt.Println(nums1)
}
