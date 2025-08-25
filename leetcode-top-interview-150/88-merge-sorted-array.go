package main

import (
	"fmt"
	"sort"
)

func merge1(nums1 []int, m int, nums2 []int, n int) []int {

	result := append(nums1[0:m], nums2[0:n]...)
	sort.Ints(result)

	return result
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	// here
	// i = last item of valid numbers of nums1
	// j = last item of nums2
	// k = last item of after merged

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	fmt.Println(nums1)
}

// func main() {
// 	fmt.Println(merge1([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
// 	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

// }
