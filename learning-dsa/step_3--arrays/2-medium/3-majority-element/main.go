package main

import "fmt"

func main() {
	nums := []int{2, 1, 2, 2, 3, 1, 3, 3, 2, 2, 2}
	fmt.Println("Majority Element: ", majorityElement(nums))
	fmt.Println("Majority Element Optimized: ", majorityElementOptimized(nums))
}

/*
Brute Force Approach - Using a Map
- Loop through the array and count the frequency of each element
- Store the frequency of each element in a map
- Find the element with the highest frequency
- Return the element with the highest frequency
*/
func majorityElement(nums []int) int {
	resMap := make(map[int]int)
	max, element := 0, 0

	for _, value := range nums {
		resMap[value]++
	}

	for key, value := range resMap {
		if value > max {
			max = value
			element = key
		}
	}

	return element
}

/*
Optimized Approach - Boyer-Moore Voting Algorithm
- Initialize a variable count to 0 and a variable candidate to 0
- Loop through the array
- If count is 0, set candidate to the current element
- If the current element is equal to the candidate, increment count by 1
- If the current element is not equal to the candidate, decrement count by 1
- Loop through the array again
- Count the frequency of the candidate
- If the frequency of the candidate is less than or equal to the length of the array divided by 2, return -1
- Return the candidate
*/
func majorityElementOptimized(nums []int) int {
	count, candidate := 0, 0

	for _, value := range nums {
		if count == 0 {
			candidate = value
		}

		if value == candidate {
			count++
		} else {
			count--
		}
	}

	cnt := 0
	for _, value := range nums {
		if value == candidate {
			cnt++
		}
	}

	if cnt <= len(nums)/2 {
		return -1
	}

	return candidate
}
