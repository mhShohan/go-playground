package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Largest element in the array is: ", largestElement(arr))
}

// brute force solution
func largestElement(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
