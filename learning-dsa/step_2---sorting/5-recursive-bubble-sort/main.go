package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort Recursive")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}

	bubbleSortRecursive(arrayOfNumbers, len(arrayOfNumbers))
	fmt.Println(arrayOfNumbers)
}

/*
*
Recursive Bubble Sort
- Base case: if n == 1, return
- Do one pass of normal Bubble Sort. This pass fixes the last element of the current subarray
- Recur for all elements except the last one
*/
func bubbleSortRecursive(arr []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSortRecursive(arr, n-1)
}
