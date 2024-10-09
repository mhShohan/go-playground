package main

import "fmt"

func main() {
	fmt.Println("---------Quick Sort --------------")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}

	result := quickSort(arrayOfNumbers, 0, len(arrayOfNumbers)-1)
	fmt.Println(result)

}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]

	return j
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		partitionIndex := partition(arr, low, high)

		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)
	}

	return arr
}
