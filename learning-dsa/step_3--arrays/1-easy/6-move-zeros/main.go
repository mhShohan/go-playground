package main

import "fmt"

func main() {
	moveZerosToEndBruteForce()

	result := moveZeros()
	fmt.Println("moveZeros Optimal Solution", result)

	numsArray := []int{10, 20, 30, 40, 50}
	res := linearSearch(numsArray, 50)
	fmt.Println("Linear Search", res)
}

// Move zeros at the end (Time Complexity: O(2n) - Space Complexity: O(n))
func moveZerosToEndBruteForce() {
	arr := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}

	temp := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp = append(temp, arr[i])
		}
	}

	copy(arr, temp)

	// for i := 0; i < len(temp); i++ {
	// 	arr[i] = temp[i]
	// }

	for i := len(temp); i < len(arr); i++ {
		arr[i] = 0
	}

	fmt.Println("Move Zeros to end (BruteForce)", arr)
}

// Optimal Solution
// Time Complexity: O(n) - Space Complexity: O(1)
func moveZeros() []int {
	arr := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}

	N := len(arr)
	j := -1

	for i := 0; i < N; i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}

	if j == -1 {
		return arr
	}

	for i := j + 1; i < N; i++ {
		if arr[i] != 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}

	return arr
}

func linearSearch(arr []int, num int) int {
	result := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			result = i
			break
		}
	}

	return result
}
