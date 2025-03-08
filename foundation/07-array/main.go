package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	// Declaring an array of fixed size 10 with string type
	var names [10]string
	names[0] = "Mehdi"
	names[1] = "Hasan"
	names[2] = "Shohan"

	fmt.Println("Names Array:", names)
	fmt.Println("Length of names array:", len(names))

	// Array with predefined values
	ages := [3]int{10, 20, 30}
	fmt.Println("Ages Array:", ages)

	// Array with inferred length
	fruits := [...]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits Array:", fruits)
	fmt.Println("Length of Fruits array:", len(fruits))

	// Iterating over an array using a for loop
	fmt.Println("\nIterating over an array:")
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index", i, ":", fruits[i])
	}

	// Iterating using range
	fmt.Println("\nUsing range to iterate:")
	for index, value := range ages {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Multi-dimensional array
	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	fmt.Println("\n2D Array (Matrix):", matrix)
}

/*
Notes:
1. **What is an Array?**
   - An array is a fixed-size, homogeneous collection of elements.
   - All elements must be of the same data type.
   - Declared with: `var arr [size]type`

2. **Array Initialization**
   - Default values are zero for numeric types and empty for strings.
   - Example: `var arr [5]int` initializes an array of 5 integers with default values.
   - You can also initialize with values: `arr := [3]int{1, 2, 3}`
   - Use `...` to let Go determine the length automatically: `arr := [...]string{"A", "B"}`

3. **Accessing Elements**
   - Use `arr[index]` to access elements.
   - Arrays are zero-indexed (`arr[0]` is the first element).

4. **Array Length**
   - Use `len(arr)` to get the number of elements.

5. **Iterating Over Arrays**
   - Use a `for` loop with an index: `for i := 0; i < len(arr); i++ {}`
   - Use `range` for better readability: `for index, value := range arr {}`

6. **Multi-Dimensional Arrays**
   - Example: `var matrix [2][2]int`
   - Access elements with `matrix[row][col]`

7. **Limitations of Arrays**
   - Fixed size (cannot be resized dynamically).
   - For dynamic collections, consider using **slices** instead.
*/
