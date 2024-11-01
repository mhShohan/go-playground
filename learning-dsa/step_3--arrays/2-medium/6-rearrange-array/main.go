package main

import "fmt"

func main() {
	arr := []int{-1, 2, 3, -4, -5, 6, 7, -8}
	uneuqalArr := []int{-1, 2, 3, -4, -5, 6, 7, -8, 9, 10}

	fmt.Println("Rearrange BruteForce: ", rearrangeBruteForce(arr))
	fmt.Println("Rearrange Optimal: ", rearrangeOptimal(arr))
	fmt.Println("Rearrange BruteForceUnequal: ", rearrangeBruteForceUnequal(uneuqalArr))
}

/*
Rearrange BruteForce
- Initialize two arrays `pos` and `neg`
- Iterate over the input array
- If the current element is less than 0, append it to the `neg` array
- Otherwise, append it to the `pos` array
- Iterate over the input array
- If the index is even, assign the element from the `pos` array
- Otherwise, assign the element from the `neg` array
- Return the input array
*/
func rearrangeBruteForce(arr []int) []int {
	pos := []int{}
	neg := []int{}

	for _, value := range arr {
		if value < 0 {
			neg = append(neg, value)
		} else {
			pos = append(pos, value)
		}
	}

	for i := 0; i < len(neg) && i < len(pos); i++ {
		arr[i*2] = pos[i]
		arr[i*2+1] = neg[i]
	}

	return arr
}

/*
Rearrange Optimal
- Initialize two variables `positiveIndex` and `negativeIndex` to 0 and 1 respectively
- Initialize an array `numbers` with the same length as the input array
- Iterate over the input array
- If the current element is less than 0, assign it to the `numbers` array at the `negativeIndex` index
- Increment the `negativeIndex` by 2
- Otherwise, assign it to the `numbers` array at the `positiveIndex` index
- Increment the `positiveIndex` by 2
- Return the `numbers` array
*/
func rearrangeOptimal(arr []int) []int {
	positiveIndex, negativeIndex := 0, 1
	numbers := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			numbers[negativeIndex] = arr[i]
			negativeIndex += 2
		} else {
			numbers[positiveIndex] = arr[i]
			positiveIndex += 2
		}
	}

	return numbers
}

/*
rearrange BruteForce Unequal
- Initialize two arrays `pos` and `neg`
- Iterate over the input array
- If the current element is less than 0, append it to the `neg` array
- Otherwise, append it to the `pos` array
- If the length of the `pos` array is greater than the length of the `neg` array
- Iterate over the `neg` array
- Assign the elements from the `pos` and `neg` arrays to the input array at even indices
- Assign the remaining elements from the `pos` array to the input array
- Otherwise
- Iterate over the `pos` array
- Assign the elements from the `pos` and `neg` arrays to the input array at even indices
- Assign the remaining elements from the `neg` array to the input array
- Return the input array
*/
func rearrangeBruteForceUnequal(arr []int) []int {
	pos := []int{}
	neg := []int{}

	for _, value := range arr {
		if value < 0 {
			neg = append(neg, value)
		} else {
			pos = append(pos, value)
		}
	}

	if len(pos) > len(neg) {
		for i := 0; i < len(neg)/2; i++ {
			arr[i*2] = pos[i]
			arr[i*2+1] = neg[i]
		}
		index := len(neg)
		for i := len(neg); i < len(pos); i++ {
			arr[index] = pos[i]
			index++
		}
	} else {
		for i := 0; i < len(pos)/2; i++ {
			arr[i*2] = pos[i]
			arr[i*2+1] = neg[i]
		}
		index := len(pos)
		for i := len(pos); i < len(neg); i++ {
			arr[index] = neg[i]
			index++
		}
	}

	return arr
}
