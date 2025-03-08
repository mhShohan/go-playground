package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	// Declare and initialize a slice
	var names = []string{"John", "Doe", "Jane"}
	fmt.Printf("Type of names slice: %T\n", names)

	// Append new elements to the slice
	names = append(names, "Shohan", "Rahman")
	fmt.Println("After append:", names)

	// Slicing operation (extracting a portion of a slice)
	names = append(names[1:3])
	fmt.Println("After slicing operation:", names)

	// Declare and initialize a numeric slice
	numbers := []int{1, 2, 3}
	fmt.Println("Original slice:", numbers)

	// Append elements to the slice
	numbers = append(numbers, 4, 5)
	fmt.Println("After append:", numbers)

	// Creating a slice using make
	nums := make([]int, 5)
	nums[0] = 121
	nums[1] = 22
	nums[2] = 332
	nums[3] = 42
	nums[4] = 51
	nums = append(nums, 61, 2323, 234)

	fmt.Println("Slice with make:", nums)
	fmt.Println("Is sorted before sorting:", sort.IntsAreSorted(nums))

	// Sorting the slice
	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)

	// Removing an element from slice based on index
	persons := []string{"John", "Doe", "Jane", "Shohan", "Rahman"}
	index := 2
	persons = append(persons[:index], persons[index+1:]...)
	fmt.Println("After removing element:", persons)

	// Copying a slice
	copySlice := make([]string, len(persons))
	copy(copySlice, persons)
	fmt.Println("Copied slice:", copySlice)

	// Iterating over a slice using range
	fmt.Println("Iterating using range:")
	for i, val := range persons {
		fmt.Println("Index:", i, "Value:", val)
	}
}

/*
Notes:
1. **What is a Slice?**
   - A slice is a dynamically-sized, flexible view into the elements of an array.
   - Unlike arrays, slices do not require a fixed size.

2. **Creating a Slice**
   - Literal declaration: `slice := []int{1, 2, 3}`
   - Using `make`: `nums := make([]int, length, capacity)`

3. **Appending Elements**
   - Use `append(slice, value)` to add elements dynamically.
   - `append(slice1, slice2...)` merges slices.

4. **Slicing Operations**
   - `slice[start:end]` extracts a portion of a slice.
   - Omitting `start` or `end` uses defaults (`0` or `len(slice)`).

5. **Removing an Element**
   - Use `append(slice[:index], slice[index+1:]...)` to remove an element.

6. **Sorting a Slice**
   - `sort.Ints(slice)` sorts an integer slice.
   - `sort.Strings(slice)` sorts a string slice.

7. **Copying Slices**
   - Use `copy(dest, src)` to copy slices efficiently.

8. **Iterating Over a Slice**
   - Use a `for` loop with `range`: `for i, val := range slice {}`

9. **Difference Between Arrays and Slices**
   - Arrays have a fixed size, while slices are dynamic.
   - Slices are backed by an underlying array, making them more flexible.
*/
