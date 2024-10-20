package main

import "fmt"

func main() {
	leftRotate([]int{1, 2, 3, 4, 5})
	leftRotateByK([]int{1, 2, 3, 4, 5, 6, 7}, 10)
	rotateRightByK([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	rotateLeftAnother([]int{1, 2, 3, 4, 5, 6, 7}, 13)
}

func leftRotate(arr []int) {
	n := len(arr)

	// Store the first element of the array
	temp := arr[0]

	// Shift the elements to the left
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}

	// Store the first element in the last position
	arr[n-1] = temp

	fmt.Println("Left Rotate 1 step", arr)
}

// left rotate by k steps
func leftRotateByK(arr []int, k int) {
	n := len(arr)
	k = k % n

	for i := 1; i <= k; i++ {
		temp := arr[0]

		for i := 1; i < n; i++ {
			arr[i-1] = arr[i]
		}

		arr[n-1] = temp
	}

	fmt.Println("Left Rotate by K step", arr)
}

/*
189. Rotate Array
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
func rotateRightByK(nums []int, k int) {
	n := len(nums)

	for i := 1; i <= k; i++ {
		temp := nums[n-1]

		for i := n - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}

		nums[0] = temp
	}

	fmt.Println("Right Rotate by k step", nums)
}

// rotate left another way - using temp array (Time complexity: O(n), Space complexity: O(k))
func rotateLeftAnother(nums []int, k int) {
	n := len(nums)
	k = k % n

	temp := []int{}
	for i := 0; i < k; i++ {
		temp = append(temp, nums[i])
	}

	for i := k; i < n; i++ {
		nums[i-k] = nums[i]
	}

	for i := 0; i < len(temp)-1; i++ {
		nums[n-k+i] = temp[i]
	}

	fmt.Println("Temp", temp, nums)
}

// optimized solution (Time complexity: O(2n), Space complexity: O(1))
func rotateOptimize(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums, 0, k-1) // reverse fast elements which rotate to the left (Time complexity: O(k))
	reverse(nums, k, n-1) // reverse the rest of the elements of array (Time complexity: O(n-k))
	reverse(nums, 0, n-1) // reverse full array (Time complexity: O(n))
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
