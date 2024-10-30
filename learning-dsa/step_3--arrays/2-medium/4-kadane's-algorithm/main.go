package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println("Brute Force Approach", maxSubArrayBruteForce(nums))
	fmt.Println("Kadane's Algorithm", maxSubArrayKadane(nums))
}

/*
Brute Force Approach: O(n^2)
- Iterate over all subarrays and calculate the sum of each subarray
- Keep track of the maximum sum
- Return the maximum sum
*/

func maxSubArrayBruteForce(nums []int) int {
	maxSum := nums[0]
	indexes := []int{0, 0} // Start and end indexes of the subarray

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
				indexes[0] = i
				indexes[1] = j
			}
		}
	}

	fmt.Println("Indexes", indexes)

	return maxSum
}

/*
Kadane's Algorithm: O(n)
- Initialize two variables `maxSum` and `currentSum` to the first element of the array
- Iterate over the array starting from the second element
- Calculate the `currentSum` as the maximum of the current element and the sum of the current element and the `currentSum`
- Update the `maxSum` as the maximum of the `maxSum` and the `currentSum`
- Return the `maxSum`
*/
func maxSubArrayKadane(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
