package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 4}
	fmt.Println("Next Permutation: ", nextPermutation(arr))
}

/*
Next Permutation
- Initialize a variable `i` to the length of the input array minus 2
- Iterate over the input array from right to left
  - If the current element is greater than the next element
  - Continue iterating
  - Otherwise
  - Break the loop

- If `i` is less than 0
  - Reverse the input arrays

- Otherwise
  - Initialize a variable `j` to the length of the input array minus 1
  - Iterate over the input array from right to left
  - If the element at index `j` is greater than the element at index `i`
  - Swap the elements at index `i` and `j`
  - Reverse the input array from index `i+1` to the end
  - Break the loop

- Return the input array
*/
func nextPermutation(arr []int) []int {
	i := len(arr) - 2

	for i >= 0 {
		if arr[i] < arr[i+1] {
			break
		}
		i--
	}

	if i < 0 {
		reverse(arr)
	} else {
		j := len(arr) - 1

		for j > i {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]

				reverse(arr[i+1:])
				break
			}
			j--
		}
	}

	return arr
}

/*
reverse
- Initialize two pointers `left` and `right` to 0 and the length of the input array minus 1
- Iterate while `left` is less than `right`
  - Swap the elements at index `left` and `right`
  - Increment `left` and decrement `right`
*/
func reverse(arr []int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
