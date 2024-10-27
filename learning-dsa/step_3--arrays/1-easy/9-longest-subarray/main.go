package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Longest subarray with given sum K(positives)
	fmt.Println("Longest subarray with given sum K(positives): ", longestSubArrayBruteForce(arr, 15))
	fmt.Println("Longest subarray with given sum K(positives): ", longestSubArrayTwoPointers(arr, 15))
}

// Longest subarray with given sum K(positives)
func longestSubArrayBruteForce(arr []int, k int) int {
	n := len(arr)
	maxLen := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == k {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}

	return maxLen
}

// tow pointers
func longestSubArrayTwoPointers(arr []int, k int) int {
	n := len(arr)
	maxLen := 0
	sum := arr[0]
	left, right := 0, 0

	for right < n {
		for left <= right && sum > k {
			sum -= arr[left]
			left++
		}

		if sum == k {
			maxLen = max(maxLen, right-left+1)
		}

		right++
		if right < n {
			sum += arr[right]
		}
	}

	return maxLen
}

// 3105. Longest Strictly Increasing or Strictly Decreasing Subarray
func longestMonotonicSubarray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxLen := 1
	increasingCount, decreasingCount := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasingCount++
			decreasingCount = 1
		} else if nums[i] < nums[i-1] {
			decreasingCount++
			increasingCount = 1
		} else {
			increasingCount, decreasingCount = 1, 1
		}

		maxLen = max(maxLen, max(increasingCount, decreasingCount))
	}

	return maxLen
}
