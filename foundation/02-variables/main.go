package main

import "fmt"

func main() {

	// Declare and initialize variables
	var username string = "MH Shohan" // Explicit declaration
	fmt.Println(username)
	fmt.Printf("Type of username: %T \n", username)

	var age int = 25 // Integer type
	fmt.Println(age)
	fmt.Printf("Type of age: %T \n", age)

	var isMarried bool = false // Boolean type
	fmt.Println(isMarried)
	fmt.Printf("Type of isMarried: %T \n", isMarried)

	// Default values of uninitialized variables
	var defaultString string // Default value: ""
	var defaultInt int       // Default value: 0
	var defaultBool bool     // Default value: false
	var defaultFloat float64 // Default value: 0.0

	fmt.Println("Default values:")
	fmt.Println("String:", defaultString)
	fmt.Println("Int:", defaultInt)
	fmt.Println("Bool:", defaultBool)
	fmt.Println("Float:", defaultFloat)

	// Short-hand declaration (Type Inference)
	name := "MH Shohan" // Go infers type as string
	gpa := 3.9254       // Go infers type as float64
	isEmployed := true  // Go infers type as bool
	fmt.Println(name, gpa, isEmployed)

	// Multiple variable declarations
	var country, city string = "Bangladesh", "Dhaka" // Declaring multiple variables
	var x, y, z int = 10, 20, 30                     // Multiple integers
	fmt.Println(country, city)
	fmt.Println(x, y, z)

	// Constants (unchangeable variables)
	const pi float64 = 3.14159
	const gravity = 9.8 // Type inferred as float64
	fmt.Println("Pi:", pi)
	fmt.Println("Gravity:", gravity)

	// Variable swapping without temporary variable
	a, b := 5, 10
	a, b = b, a
	fmt.Println("Swapped values:", a, b)
}

/*
Notes:
1. `var` is used for explicit declaration and initialization.
2. Variables have default values: string (""), int (0), bool (false), float (0.0).
3. `:=` is used for shorthand declaration with type inference.
4. Constants (`const`) cannot be changed after assignment.
5. Multiple variables can be declared in a single line.
6. Variables can be swapped without a temp variable using `a, b = b, a`.
7. The `fmt.Printf("%T", varName)` is used to check the type of a variable.
*/
