package main

import "fmt"

func main() {
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 12, 25, 3, 24}
	fmt.Println("---------------------Selection Sort--------------------")

	result := selectionSort(arrayOfNumbers)
	fmt.Println(result)
}

func selectionSort(arr []int) []int {
	count := 0 // for iteration count
	for i := 0; i <= len(arr)-2; i++ {
		min := i
		for j := i; j <= len(arr)-1; j++ {
			if arr[j] < arr[min] {
				min = j
			}

			count++ // for iteration count
		}

		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}

	fmt.Println("iteration count", count)
	return arr
}
