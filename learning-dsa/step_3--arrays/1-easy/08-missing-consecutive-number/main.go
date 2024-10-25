package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 8, 9, 10}
	// arr := []int{3, 0, 1}
	missingNumber := findMissingNumberXor(arr)
	println(missingNumber)

	consArr := []int{1, 1, 0, 1, 1, 1, 0, 0, 1}
	fmt.Println("Ones", maxConsecutiveOne(consArr))
}

func findMissingNumberBruteForce(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return -1
}

func findMissingNumberSum(arr []int) int {
	n := len(arr)
	sum := (n) * (n + 1) / 2 // sum of n natural numbers

	for i := 0; i < n; i++ {
		sum -= arr[i] // subtract all the elements in the array
	}

	return sum
}

func findMissingNumberXor(arr []int) int {
	N := len(arr)
	x1, x2 := 0, 0

	for i := 0; i < N; i++ {
		x2 = x2 ^ arr[i]
		x1 = x1 ^ (i + 1)
	}

	return x1 ^ x2
}

func maxConsecutiveOne(arr []int) int {
	count := 0
	maxCount := 0

	for _, value := range arr {
		if value == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}
