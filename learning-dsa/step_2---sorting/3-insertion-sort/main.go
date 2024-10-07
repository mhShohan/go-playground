package main

import "fmt"

func main() {
	fmt.Println("----------Insertion sort-------------------")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}

	result := inserTionSort(arrayOfNumbers)
	fmt.Println(result)
}

func inserTionSort(arr []int) []int {
	count := 0 // for iteration count
	for i := 0; i <= len(arr)-1; i++ {

		j := i

		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--

			count++ // for iteration count
		}
	}

	fmt.Println("iteration count:", count)
	return arr
}
