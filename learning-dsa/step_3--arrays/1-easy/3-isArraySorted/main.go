package main

import "fmt"

func main() {

	fmt.Println("Is array sorted: ", isArraySorted([]int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println("Is array sorted: ", isArraySorted([]int{3, 4, 5, 1, 2}))

}

func isArraySorted(arr []int) bool {
	isSorted := true
	element := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] >= element {
			element = arr[i]
		} else {
			isSorted = false
			break
		}
	}

	return isSorted
}

// 1752. Check if Array Is Sorted and Rotated --- Leetcode
func check(nums []int) bool {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			count++
		}
	}

	return count <= 1
}
