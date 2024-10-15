package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 9, 4, 5, 6, 7, 8, 8, 9, 9}

	fmt.Println("Second largest element (brute force solution): ", secondLargestElementBruteForce(arr))
	fmt.Println("Second largest element (optimized solution): ", secondLargestElement(arr))

	fmt.Println("Second smallest element: ", secondSmallestElement(arr))
}

// brute force solution using sort
func secondLargestElementBruteForce(arr []int) int {
	// O(nlogn) time complexity for sorting the array - try to implement sorting by owns
	sort.Ints(arr)
	max := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] < max {
			max = arr[i]
			break
		}
	}

	return max
}

// optimized solution
func secondLargestElement(arr []int) int {
	max := arr[0]
	secondMax := math.MinInt32

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		} else if arr[i] > secondMax && max > arr[i] {
			secondMax = arr[i]
		}
	}

	return secondMax
}

// second smallest element
func secondSmallestElement(arr []int) int {
	min := arr[0]
	secondMin := math.MaxInt32

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			secondMin = min
			min = arr[i]
		} else if arr[i] < secondMin && min < arr[i] {
			secondMin = arr[i]
		}
	}

	return secondMin
}
