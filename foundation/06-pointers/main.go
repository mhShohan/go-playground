package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// Declaring a pointer without initialization (default is nil)
	var ptr *int
	fmt.Println("Value of ptr (uninitialized pointer):", ptr) // Output: nil

	// Initializing an integer variable
	num := 10

	// Creating a pointer that stores the address of num
	var ptr2 = &num

	fmt.Println("Address stored in ptr2 (Address of num):", ptr2)
	fmt.Println("Value at the address stored in ptr2 (*ptr2):", *ptr2) // Dereferencing pointer

	// Modifying the value through the pointer
	*ptr2 = 20
	fmt.Println("Updated value of num (through pointer):", num) // num is now 20

	// Using pointers with functions
	fmt.Println("\nUsing pointers with functions")
	number := 50
	fmt.Println("Before function call, number:", number)
	modifyValue(&number) // Passing address to function
	fmt.Println("After function call, number:", number)

	// Pointer to a struct
	fmt.Println("\nPointers with Structs")
	type Person struct {
		name string
		age  int
	}

	person := Person{"John", 30}
	personPtr := &person
	fmt.Println("Person Name (via pointer):", personPtr.name) // Accessing fields via pointer
	fmt.Println("Person Age (via pointer):", personPtr.age)

	// Modifying struct fields through pointer
	personPtr.age = 35
	fmt.Println("Updated Person Age (via pointer):", person.age)
}

// Function that modifies a value using a pointer
func modifyValue(val *int) {
	*val += 10 // Increment value by 10
}

/*
Notes:
1. **What are Pointers?**
   - A pointer is a variable that stores the memory address of another variable.
   - In Go, pointers are declared using the `*` symbol.
   - Example: `var ptr *int` (declares a pointer to an integer).

2. **Dereferencing a Pointer**
   - Using `*ptr`, we can access or modify the value stored at the address the pointer points to.
   - Example: `*ptr = 20` modifies the original variable.

3. **Passing Pointers to Functions**
   - Instead of passing copies, we can pass memory addresses to functions, allowing modifications.
   - Example: `modifyValue(&number)` passes the address of `number` to the function.

4. **Pointers to Structs**
   - We can use pointers to access and modify struct fields efficiently.
   - Example: `personPtr.age = 35` modifies `person.age` directly.

5. **Why Use Pointers?**
   - Efficient memory usage (avoids copying large values).
   - Allows functions to modify original values.
   - Enables dynamic memory management.

Remember:
- Go does not support pointer arithmetic (unlike C/C++).
- Use `&` to get the address of a variable.
- Use `*` to dereference a pointer and get its value.
*/
