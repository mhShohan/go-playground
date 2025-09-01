package main

import "fmt"

/**
 * Rotate Right Brute Force
 * Steps:
 *   - Store the last element
 *   - Shift all elements to the right
 *   - Place the stored element at the beginning
 */
func rotateRightBruteForce(nums []int, k int) {
	lastIndex := len(nums) - 1
	i := 0

	for i < k {
		temp := nums[lastIndex]
		for j := lastIndex; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = temp
		i++
	}

	fmt.Println(nums)
}

/**
 * Rotate Left Brute Force
 * Steps:
 *   - Store the first element
 *   - Shift all elements to the left
 *   - Place the stored element at the end
 */
func rotateLeftBruteForce(nums []int, k int) {
	rotations := k % len(nums)
	lastIndex := len(nums) - 1

	for i := 0; i < rotations; i++ {
		first := nums[0]
		for j := 0; j < lastIndex; j++ {
			nums[j] = nums[j+1]
		}
		nums[lastIndex] = first
	}

	fmt.Println(nums)
}

/*
	Rotate Left Optimized
	Steps:
		- Reverse the first k elements - Time O(k)
		- Reverse the remaining elements - Time O(n-k)
		- Reverse the entire array - Time O(n)
*/

func rotateLeftOptimized(nums []int, k int) {
	rotations := k % len(nums)
	reverse(nums, 0, rotations-1)
	reverse(nums, rotations, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// Reverse the elements in the array from start to end
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/*
	Rotate Right Optimized
	Steps:
		- Reverse the last k elements - Time O(k)
		- Reverse the remaining elements - Time O(n-k)
		- Reverse the entire array - Time O(n)
*/

func rotateRightOptimized(nums []int, k int) {
	rotations := k % len(nums)
	reverse(nums, len(nums)-rotations, len(nums)-1)
	reverse(nums, 0, len(nums)-rotations-1)
	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// func main() {
// 	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
// 	fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))
// }
