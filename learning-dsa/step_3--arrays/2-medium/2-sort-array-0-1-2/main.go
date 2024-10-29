package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColorsBruteForce(nums)
	sortColorsOptimal(nums)
}

// 75. Sort Colors (Medium)
// Sort an array of 0's 1's and 2's
/*
Brute Force Solution:
- Create a map to store the frequency of each element
- Traverse the array and store the frequency of each element in the map
- Traverse the map and update the array with the frequency of each element
- Time Complexity: O(2n)
- Space Complexity: O(n)
*/
func sortColorsBruteForce(nums []int) {
	resMap := make(map[int]int)

	for _, value := range nums {
		resMap[value]++
	}

	index := 0
	for i := 0; i < 3; i++ {
		for resMap[i] > 0 {
			nums[index] = i
			resMap[i]--
			index++
		}
	}

	fmt.Println("Sorted Array: ", nums)
}

/*
Optimal Solution:
- Dutch National Flag Algorithm
- Three pointers: low, mid, high
- low: 0, mid: 0, high: n-1
- Traverse the array from start to end
- If nums[mid] == 0, swap nums[low] and nums[mid], increment low and mid
- If nums[mid] == 1, increment mid
- If nums[mid] == 2, swap nums[mid] and nums[high], decrement high
- Continue until mid <= high
- Time Complexity: O(n)
- Space Complexity: O(1)
*/

func sortColorsOptimal(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}

	fmt.Println("Sorted Array: ", nums)
}
