# Types of Function

- Standard or named function
- Anonymous function
- Function expression or Assign function in variable
- Higher order function or first class function
- Callback function
- Variadic Function
- Init function - we don not call this function, computer call this function
- Closure
- Defer function - last in first out
- Receiver function
- IIFE - Immediately Invoked Function Expression

## Now let's walk through each type of function in Go:

## 1. Standard or Named Function

A basic function with a defined name that can be called throughout your code.

```go
func greet(name string) string {
    return "Hello, " + name + "!"
}
```

This is the most common type of function that you define once and can call anywhere in your program.

## 2. Anonymous Function

A function without a name, defined inline where it's needed.

```go
anonymous := func(x, y int) int {
    return x + y
}
fmt.Println(anonymous(5, 3))  // Outputs: 8
```

Anonymous functions are useful when you need a quick function for a specific task and don't want to define it separately.

## 3. Function Expression (Assigned to Variable)

Similar to anonymous functions but the focus is on storing the function in a variable.

```go
addFunc := func(a, b int) int {
    return a + b
}
fmt.Println(addFunc(10, 20))  // Outputs: 30
```

This allows you to pass functions around like any other value.

## 4. Higher-Order Function (First-Class Function)

A function that takes another function as an argument or returns a function.

```go
func processNumbers(numbers []int, processor func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = processor(num)
    }
    return result
}

// Using it:
doubled := processNumbers([]int{1, 2, 3, 4, 5}, func(n int) int {
    return n * 2
})
```

Higher-order functions enable powerful abstractions and flexible code design.

## 5. Callback Function

A function passed to another function as an argument, to be called later.

```go
// The processor function in the processNumbers example is a callback
```

Callbacks are commonly used for event handling, asynchronous operations, or customizing behavior.

## 6. Variadic Function

A function that accepts a variable number of arguments.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Can be called with any number of arguments
fmt.Println(sum(1, 2, 3))         // Outputs: 6
fmt.Println(sum(5, 10, 15, 20))   // Outputs: 50
```

The `...` syntax allows you to process any number of arguments of the same type.

## 7. Init Function

Special function that's automatically called before `main()` executes.

```go
func init() {
    fmt.Println("Init function: I run before main")
}
```

You can have multiple `init()` functions in a program; they'll all run automatically during startup.

## 8. Closure

A function that "remembers" the environment in which it was created.

```go
func createGreeter(greeting string) func(string) string {
    return func(name string) string {
        return greeting + ", " + name + "!"
    }
}

// Using closures:
spanishGreeter := createGreeter("Hola")
fmt.Println(spanishGreeter("Carlos"))  // Outputs: Hola, Carlos!
```

Closures "capture" variables from their surrounding scope, maintaining access even after the outer function returns.

## 9. Defer Function

A function call that's executed after the surrounding function returns.

```go
func main() {
    defer fmt.Println("This runs last")
    defer fmt.Println("This runs second")
    fmt.Println("This runs first")
}
```

Defer functions run in LIFO (Last In, First Out) order - great for cleanup tasks like closing files.

## 10. Receiver Function (Method)

A function attached to a type, allowing object-oriented style programming.

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Introduce() {
    fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

// Using it:
person := Person{Name: "John", Age: 30}
person.Introduce()
```

Methods let you associate behavior with data types, enabling cleaner, more intuitive code.

## 11. IIFE (Immediately Invoked Function Expression)

A function that's defined and executed immediately.

```go
func() {
    message := "I'm executed immediately!"
    fmt.Println(message)
}() // The parentheses at the end execute the function
```

IIFE is useful for one-time operations, initialization, or creating private scopes.

The full code example in the artifact demonstrates each of these function types working together in a complete Go program.
