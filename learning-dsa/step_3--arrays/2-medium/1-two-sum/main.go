package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 6, 5, 8, 11}

	fmt.Println("Brute Force solution:", twoSumBruteForce(arr, 14))
	fmt.Println("Better solution using map:", twoSumBetterSolutionMap(arr, 14))

	twoSumUsingTwoPointer(arr, 14)
}

func twoSumBruteForce(arr []int, num int) string {
	result := "No"

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == num {
				result = "Yes"
			}
		}
	}

	return result
}

func twoSumBetterSolutionMap(arr []int, target int) string {
	resultMap := make(map[int]int)
	result := "No"
	resultIndexes := []int{0, 0}

	for i := 0; i < len(arr); i++ {
		needed := target - arr[i]

		if _, ok := resultMap[needed]; !ok {
			resultMap[arr[i]] = i
		} else {
			resultIndexes[0] = resultMap[needed]
			resultIndexes[1] = i
			result = "Yes"
		}
	}

	fmt.Println(resultIndexes)

	return result
}

func twoSumUsingTwoPointer(arr []int, target int) {
	left, right := 0, len(arr)-1
	result := "no"
	sort.Ints(arr)

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			result = "yes"
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	fmt.Println("using two pointer", result)
}
