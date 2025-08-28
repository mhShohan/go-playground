package main

import (
	"fmt"
)

func removeDuplicatesII(nums []int) int {
	k := 2

	for i := k; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
		fmt.Println("iteration-", i, nums)
	}

	fmt.Println(nums[0:k])

	return k
}

// func main() {
// 	fmt.Println("resutl", removeDuplicatesII([]int{1, 1, 1, 2, 2, 3}))
// 	fmt.Println("resutl", removeDuplicatesII([]int{0, 0, 0, 1, 1, 1, 1, 2, 3, 3}))
// }
