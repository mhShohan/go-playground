package main

import "fmt"

func main() {
	fmt.Println("---------------Bubble Sort-------------------")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}
	sortedArrayOfNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	result := bubbleSort(arrayOfNumbers)
	fmt.Println(result)

	result2 := bubbleSort(sortedArrayOfNumber)
	fmt.Println(result2)
}

func bubbleSort(arr []int) []int {
	count := 0 // for iteration count

	for i := len(arr) - 1; i >= 1; i-- {
		isSwap := false // for optimization , it will run O(n) for best case or sorted array

		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				isSwap = true
			}
			count++ // for iteration count
		}

		if !isSwap {
			break
		}

	}

	fmt.Println("iteration count", count)
	return arr
}

/*
Explanation: to sorting [5,4,7,2]
iteration 1 - [5,4,7,2] => [4,5,7,2] => [4,5,2,7]
iteration 2 - [4,5,2,7] => [4,2,5,7]
iteration 3 - [4,2,5,7] => [2,4,6,7]
*/
