package main

import "fmt"

func main() {
	numsArr := []int{1, 2, 3, 3, 4, 4, 5, 5, 5, 6, 7, 8, 8, 9, 9}

	result := removeDuplicates(numsArr)
	fmt.Println("Unique element count:", result)

	removeDuplicatesFromUnsortedArray()
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	fmt.Println("Unique nums array:", nums[:i+1])
	return i + 1
}

func removeDuplicatesFromUnsortedArray() {
	arr := []int{1, 7, 6, 9, 1, 11, 2, 3, 3, 4, 4, 5, 5, 5, 6, 7, 8, 8, 9, 9}
	// Create a map to store the frequency of each element
	freqMap := make(map[int]int)

	// Loop through the array and store the frequency of each element
	for _, num := range arr {
		freqMap[num]++
	}

	// Loop through the array and remove the duplicates
	i := 0
	for j := 0; j < len(arr); j++ {
		if freqMap[arr[j]] > 0 {
			arr[i] = arr[j]
			i++
			freqMap[arr[j]] = 0
		}
	}

	fmt.Println(freqMap)

	fmt.Println("Updated arr:", arr[:i])
}
