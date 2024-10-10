package main

import "fmt"

func main() {
	fmt.Println("Insertion Sort Recursive")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}
	anotherArrayOfNumbers := []int{5, 32, 1, 24, 20, 13, 54, 75, 16, 73, 88, 9, 10}

	insertionSortRecursive(arrayOfNumbers, len(arrayOfNumbers))
	fmt.Println(arrayOfNumbers)

	insertion_sort(anotherArrayOfNumbers, 1, len(anotherArrayOfNumbers))
	fmt.Println(anotherArrayOfNumbers)
}

/*
*
Recursive Insertion Sort
- Base case: if n <= 1, return
- Recur for n-1 elements
- Insert last element at its correct position in sorted array of n-1 elements
*/
func insertionSortRecursive(arr []int, n int) {
	if n <= 1 {
		return
	}

	insertionSortRecursive(arr, n-1)

	last := arr[n-1]
	j := n - 2

	for j >= 0 && arr[j] > last {
		arr[j+1] = arr[j]
		j--
	}

	arr[j+1] = last
}

// another
func insertion_sort(arr []int, i, n int) {
	if n == 1 {
		return
	}

	j := i
	for j > 0 && arr[j-1] > arr[j] {
		arr[j-1], arr[j] = arr[j], arr[j-1]
		j--
	}

	insertion_sort(arr, i+1, n-1)
}
