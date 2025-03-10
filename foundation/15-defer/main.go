package main

import "fmt"

func main() {
	fmt.Println("--------Defer---------")

	// Defer follows LIFO (Last In, First Out)
	defer fmt.Println("1")
	defer fmt.Println("2")

	fmt.Println("3")
	fmt.Println("4")

	defer fmt.Println("5")

	fmt.Println("6")

	testDefer()
}

// Function demonstrating defer with multiple statements
func testDefer() {
	fmt.Println("--------Defer in Function---------")
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	fmt.Println("D")
	defer fmt.Println("E")
}

/*
Notes:
=========
- **Defer in Go**
  - `defer` postpones the execution of a statement until the surrounding function returns.
  - Multiple `defer` statements are executed in **LIFO (Last In, First Out)** order.
  - Useful for resource cleanup, such as closing files or database connections.

- **Key Execution Flow**
  - Regular statements execute in order.
  - Deferred statements are stored and executed in reverse order when the function exits.

Key Takeaways:
- Use `defer` for cleanup operations (e.g., closing files, unlocking mutexes).
- Remember the LIFO order of execution.
- Defer can be used multiple times in a function.
*/
