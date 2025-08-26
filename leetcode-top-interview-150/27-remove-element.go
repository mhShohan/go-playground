package main

import (
	"fmt"
)

// 26. Remove Element
func removeElement(nums []int, val int) int {

	// if nums is an empty array the return 0
	if len(nums) == 0 {
		return 0
	}

	k := 0 // first pointer
	i := 0 // second pointer
	for i < len(nums) {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}

		i++
	}

	fmt.Println(nums)

	return k
}

// func main() {
// 	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
// 	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
// }
