package main

import "fmt"

func main() {
	fmt.Println("--- Functions ---")

	// Example 1: Basic Function Call
	result := sum(10, 20)
	fmt.Println("Result:", result)

	// Example 2: Variadic Function
	result = sumOfNumbers(10, 20, 30, 40, 50)
	fmt.Println("Result:", result)

	// Example 3: Function Returning Multiple Values
	sum, sub := sumAndSubtract(100, 20)
	fmt.Println("Sum:", sum, "Subtract:", sub)

	// Example 4: Anonymous Function (Immediately Invoked Function)
	func() {
		fmt.Println("IFII (Immediately Invoked Function)")
	}()

	// Example 5: Function as a Value
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiplication Result:", multiply(5, 4))

	// Example 6: Higher-Order Function (Function Returning Another Function)
	doubler := createMultiplier(2)
	fmt.Println("Doubled 10:", doubler(10))
}

// Basic Function
func sum(a int, b int) int {
	return a + b
}

// Variadic Function
func sumOfNumbers(a ...int) int {
	total := 0
	for _, value := range a {
		total += value
	}
	return total
}

// Function Returning Multiple Values
func sumAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// Function Returning Another Function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

/*
Notes:
=========
- **Functions in Go**
  - Defined using `func functionName(params) returnType {}` syntax.
  - Supports multiple return values (`func foo() (int, string)`).
  - Variadic functions allow passing multiple arguments (`func foo(nums ...int)`).
  - Anonymous functions can be defined inline and used immediately.
  - Functions can be assigned to variables and passed as arguments.
  - Higher-order functions can return other functions, allowing for closures.

Key Takeaways:
- Prefer variadic functions when handling a variable number of arguments.
- Use multiple return values instead of modifying function parameters.
- Anonymous functions are useful for quick, one-time operations.
- Higher-order functions enable more flexible and reusable code.
*/
