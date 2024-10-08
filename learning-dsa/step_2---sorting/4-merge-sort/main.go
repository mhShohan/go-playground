package main

import "fmt"

func main() {
	fmt.Println("-------------Merge Sort------------------")
	arrayOfNumbers := []int{11, 7, 14, 23, 37, 32, 44, 12, 25, 3, 24}

	result := mergeSort(arrayOfNumbers, 0, len(arrayOfNumbers)-1)
	fmt.Println(result)
}

func merge(arr []int, low, mid, high int) {
	temp := []int{}

	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

func mergeSort(arr []int, low int, high int) []int {
	if low == high {
		return arr
	}

	var mid int = (low + high) / 2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)

	merge(arr, low, mid, high)

	return arr
}
