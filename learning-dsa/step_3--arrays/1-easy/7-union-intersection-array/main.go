package main

import (
	"fmt"
	"sort"
)

func main() {

	arrOne := []int{1, 2, 4, 4, 5, 6}
	arrTwo := []int{1, 3, 4, 4, 6, 7, 8}

	result := unionOfTwoSortedArrayBruteForce(arrOne, arrTwo)
	fmt.Println("Brute Force Solution: ", result)

	result = unionOfTwoSortedArray(arrOne, arrTwo)
	fmt.Println("Optimal Solution Solution: ", result)

	intersectionOfTwoSortedArray(arrOne, arrTwo)
}

func unionOfTwoSortedArrayBruteForce(arrOne, arrTwo []int) []int {
	uniqueValues := make(map[int]int)

	for i := 0; i < len(arrOne); i++ {
		if _, exists := uniqueValues[arrOne[i]]; !exists {
			uniqueValues[arrOne[i]] = arrOne[i]
		}
	}

	for i := 0; i < len(arrTwo); i++ {
		if _, exists := uniqueValues[arrTwo[i]]; !exists {
			uniqueValues[arrTwo[i]] = arrTwo[i]
		}
	}

	uniqueValuesArray := []int{}

	for _, value := range uniqueValues {
		uniqueValuesArray = append(uniqueValuesArray, value)
	}

	// sort the array
	sort.Ints(uniqueValuesArray)

	return uniqueValuesArray
}

func unionOfTwoSortedArray(arrOne, arrTwo []int) []int {
	union := []int{}

	i, j := 0, 0
	n1, n2 := len(arrOne), len(arrTwo)

	for i < n1 && j < n2 {
		if arrOne[i] <= arrTwo[j] {
			// check the element is not in the union array (to avoid duplicates) then add it
			if len(union) == 0 || union[len(union)-1] != arrOne[i] {
				union = append(union, arrOne[i])
			}
			i++
		} else {
			// check the element is not in the union array (to avoid duplicates) then add it
			if len(union) == 0 || union[len(union)-1] != arrTwo[j] {
				union = append(union, arrTwo[j])
			}
			j++
		}
	}

	// add the remaining elements of arrOne (if any)
	for i < n1 {
		if len(union) == 0 || union[len(union)-1] != arrOne[i] {
			union = append(union, arrOne[i])
		}
		i++
	}

	// add the remaining elements of arrTwo (if any)
	for j < n2 {
		if len(union) == 0 || union[len(union)-1] != arrTwo[j] {
			union = append(union, arrTwo[j])
		}
		j++
	}

	return union
}

func intersectionOfTwoSortedArray(arrOne, arrTwo []int) {
	intersection := []int{}

	i, j := 0, 0
	n1, n2 := len(arrOne), len(arrTwo)

	for i < n1 && j < n2 {
		if arrOne[i] < arrTwo[j] {
			i++
		} else if arrOne[i] > arrTwo[j] {
			j++
		} else {
			intersection = append(intersection, arrOne[i])
			i++
			j++
		}
	}

	fmt.Println("intersection Two Sorted Array", intersection)
}
