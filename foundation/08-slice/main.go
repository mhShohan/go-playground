package main

import (
	"fmt"
	"sort"
)

func main() {
	var names = []string{"John", "Doe", "Jane"}
	fmt.Printf("Types of slices: %T\n", names)

	names = append(names, "Shohan", "Rahman")
	names = append(names[1:3])
	fmt.Println("After append and slice:", names)

	// Declare and initialize a slice
	numbers := []int{1, 2, 3}
	fmt.Println("Slice:", numbers)

	// Append elements to the slice
	numbers = append(numbers, 4, 5)
	fmt.Println("After append:", numbers)

	// Create a slice using make
	nums := make([]int, 5)

	nums[0] = 121
	nums[1] = 22
	nums[2] = 332
	nums[3] = 42
	nums[4] = 51

	nums = append(nums, 61, 2323, 234)

	fmt.Println("Slice with make:", nums)
	fmt.Println("Is sorted:", sort.IntsAreSorted(nums))

	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)

	// remove element from slice based on index
	persons := []string{"John", "Doe", "Jane", "Shohan", "Rahman"}
	index := 2
	persons = append(persons[:index], persons[index+1:]...)

	fmt.Println("After remove:", persons)
}

// Slice - A slice in Go is a dynamically-sized, flexible view into the elements of an array. Slices are more powerful and convenient than arrays because they provide a way to work with sequences of data without having to specify the size upfront.
