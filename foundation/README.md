# Basic Golang Learning Guide

This guide will help you learn Go (Golang) from fundamentals to advanced concepts, with TypeScript comparisons for easier understanding.

## What is Go?

Go is a statically typed, compiled language designed by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson. It focuses on simplicity, efficiency, and built-in concurrency support.

**Explanation:** Go was created to address the challenges of building scalable network services and large codebases. It combines the performance of compiled languages like C++ with the simplicity and readability of Python. Unlike TypeScript, which builds on JavaScript, Go was designed from scratch with modern computing challenges in mind.

## Why Go?

- **Performance**: Compiles to machine code (unlike TypeScript → JavaScript)
- **Concurrency**: Native support through goroutines and channels
- **Simplicity**: Minimal syntax with few keywords
- **Strong Standard Library**: Robust built-in packages
- **Fast Compilation**: Quick build times

**Explanation:** As a TypeScript developer, you'll find Go offers superior performance since it compiles directly to machine code instead of being interpreted like JavaScript. Go's approach to concurrency is fundamentally different and more efficient than JavaScript's async/await pattern, making it ideal for high-throughput services.

## Go vs TypeScript: Key Differences

| Feature               | Go                                     | TypeScript                      |
| --------------------- | -------------------------------------- | ------------------------------- |
| Type System           | Static, strict                         | Static, configurable strictness |
| Execution             | Compiled to binary                     | Transpiled to JavaScript        |
| Concurrency           | Native (goroutines)                    | Async/await, Promises           |
| Object Model          | No classes (structs + interfaces)      | Classes, interfaces             |
| Dependency Management | Go modules                             | npm/yarn                        |
| Error Handling        | Explicit returns                       | Exceptions                      |
| Paradigm              | Primarily procedural with OOP features | Object-oriented, functional     |

**Explanation:** Go takes a different approach to many programming concepts you're familiar with in TypeScript. For instance, Go doesn't use exceptions for error handling; instead, it returns errors as values that must be explicitly checked. This leads to more explicit code with fewer unexpected behaviors.

## Getting Started

1. **Install Go**: Download from [golang.org](https://golang.org/dl/)
2. **Verify installation**: Run `go version` in terminal
3. **Setup workspace**: Create a directory for your Go projects

**Explanation:** Unlike Node.js projects where dependencies are installed in each project folder, Go traditionally used a single workspace for all Go code. However, with Go modules (introduced in Go 1.11), you can create projects anywhere on your system, similar to npm/yarn projects.

## Basic Go Concepts

### Variables and Types

```go
// Variable declaration
var name string = "John"
age := 30  // Type inference with :=

// Basic types
var isActive bool = true
var count int = 42
var price float64 = 19.99
var message string = "Hello Go"

// Constants
const PI = 3.14159

// Multiple variable declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    height    int    = 180
)
```

**Explanation:** Go offers two ways to declare variables: the `var` keyword and the short declaration operator (`:=`). Unlike TypeScript, Go variables must have a value and cannot be `undefined`. Go's type system is also more rigid than TypeScript's, with fewer implicit type conversions.

### Functions

```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // Returns named values automatically
}

// Using the function with error handling
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

**Explanation:** Functions in Go can return multiple values, which is ideal for returning both results and errors. This differs from TypeScript where you might use Promises or custom objects to return multiple values. Go also supports variadic functions (like JavaScript's rest parameters) and function closures.

### Control Structures

```go
// If statement
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 10 {
    fmt.Println("x equals 10")
} else {
    fmt.Println("x is less than 10")
}

// If with a short statement
if value, err := someFunction(); err == nil {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Error:", err)
}

// For loop (Go has only one loop keyword: for)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// For as a while loop
counter := 0
for counter < 5 {
    fmt.Println(counter)
    counter++
}

// Infinite loop
for {
    // Do something forever
    break // Exit loop
}

// Range loop (for arrays, slices, maps, channels)
fruits := []string{"apple", "banana", "orange"}
for index, fruit := range fruits {
    fmt.Printf("%d: %s\n", index, fruit)
}

// Switch statement
switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("End of work week")
default:
    fmt.Println("Regular day")
}

// Switch with no expression (alternative to if-else chains)
switch {
case age < 13:
    fmt.Println("Child")
case age < 20:
    fmt.Println("Teenager")
default:
    fmt.Println("Adult")
}
```

**Explanation:** Go's control structures are similar to TypeScript but with some important differences. Go doesn't have a traditional while loop; instead, you use the `for` keyword for all looping constructs. The `range` keyword simplifies iteration over collections. Unlike TypeScript/JavaScript, Go's `switch` statements don't fall through to the next case by default, so `break` statements are unnecessary.

### Data Structures

```go
// Arrays (fixed size)
var numbers [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(len(numbers)) // Length: 5

// Slices (dynamic size)
fruits := []string{"apple", "banana", "orange"}
fruits = append(fruits, "mango")
newFruits := fruits[1:3] // Create slice from index 1 to 2

// Make function to create slices with capacity
numbers := make([]int, 5, 10) // Length 5, capacity 10

// Maps (like objects/dictionaries)
userInfo := map[string]string{
    "name": "John",
    "email": "john@example.com",
}

// Check if a key exists
if value, exists := userInfo["phone"]; exists {
    fmt.Println("Phone:", value)
} else {
    fmt.Println("No phone number found")
}

// Delete a key
delete(userInfo, "email")
```

**Explanation:** Go makes a clear distinction between fixed-size arrays and dynamic slices. Slices are more commonly used and are similar to JavaScript arrays but with better performance characteristics. Maps in Go are similar to JavaScript objects/dictionaries, but Go requires all keys in a map to be of the same type, and all values to be of the same type.

### Structs (Similar to TypeScript interfaces/classes)

```go
// Define a struct
type User struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
    IsActive  bool
    CreatedAt time.Time
}

// Create an instance
user := User{
    ID:        1,
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john@example.com",
    IsActive:  true,
    CreatedAt: time.Now(),
}

// Short variable declaration
anotherUser := User{ID: 2, FirstName: "Jane"}

// Access fields
fmt.Println(user.Email)

// Anonymous structs
point := struct {
    X, Y int
}{
    X: 10,
    Y: 20,
}

// Embedded structs (composition)
type Address struct {
    Street  string
    City    string
    Country string
}

type Employee struct {
    User    // Embedded struct
    Address // Embedded struct
    Salary  float64
}
```

**Explanation:** Structs in Go are similar to TypeScript interfaces but also serve as data containers like classes. Go doesn't have inheritance; instead, it uses composition through embedded structs. This encourages a different design approach focusing on behavior (interfaces) rather than type hierarchies.

### Methods

```go
// Methods on structs
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver (like class methods in TypeScript)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver for modifying the struct
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()  // Call the method
rect.Scale(2)        // Modify the struct
fmt.Println(rect.Width) // 20
```

**Explanation:** Methods in Go are functions with a receiver parameter. The receiver can be a value or a pointer. Using pointer receivers allows methods to modify the receiver (like `this` in TypeScript). Unlike TypeScript classes, Go methods can be defined separately from the struct definition, even in different packages.

### Interfaces

```go
// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Multiple types can implement an interface
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle also implements Shape
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Use interface as parameter type
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

// Empty interface can hold any value (similar to 'any' in TypeScript)
func PrintAny(v interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Type assertions
func process(i interface{}) {
    if value, ok := i.(string); ok {
        fmt.Println("String:", value)
    } else if value, ok := i.(int); ok {
        fmt.Println("Integer:", value)
    } else {
        fmt.Println("Unknown type")
    }
}

// Type switches
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

**Explanation:** Interfaces in Go are implicitly implemented (no `implements` keyword like in TypeScript). Any type that has all the methods defined in an interface automatically implements that interface. This design promotes loose coupling between packages. Empty interfaces (`interface{}`) are similar to TypeScript's `any` type, allowing for functions that accept any type of value.

### Concurrency

```go
// Goroutines (lightweight threads)
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

// Start a goroutine
go printNumbers()  // Non-blocking
fmt.Println("Started")  // This executes immediately

// Channels for communication between goroutines
ch := make(chan string)

go func() {
    ch <- "Hello from goroutine"  // Send to channel
}()

message := <-ch  // Receive from channel (blocks until data arrives)
fmt.Println(message)

// Buffered channels
bufferedCh := make(chan int, 3)
bufferedCh <- 1
bufferedCh <- 2
bufferedCh <- 3
// The channel can hold 3 values without blocking

// Select statement (like switch for channels)
select {
case msg1 := <-channel1:
    fmt.Println("Received from channel 1:", msg1)
case msg2 := <-channel2:
    fmt.Println("Received from channel 2:", msg2)
case channel3 <- "hello":
    fmt.Println("Sent to channel 3")
default:
    fmt.Println("No channel operations ready")
}

// Closing channels and ranging over them
func producer(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch) // Signal that no more values will be sent
}

ch := make(chan int)
go producer(ch)

// Range over channel until closed
for num := range ch {
    fmt.Println(num)
}
```

**Explanation:** Concurrency is where Go truly shines compared to TypeScript. Goroutines are extremely lightweight threads (you can create thousands without issue) that are managed by the Go runtime. Channels provide a safe way for goroutines to communicate without shared memory. This is a fundamentally different approach from JavaScript's event loop and async/await pattern, offering better performance for parallel operations.

### Error Handling

```go
// Creating errors
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Custom error types
type ValidationError struct {
    Field string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("Validation error on field %s: %s", e.Field, e.Message)
}

// Using errors.Is and errors.As (Go 1.13+)
if errors.Is(err, os.ErrNotExist) {
    // Handle specific error
}

var validationErr ValidationError
if errors.As(err, &validationErr) {
    // Handle validation error
    fmt.Println(validationErr.Field)
}
```

**Explanation:** Go uses explicit error checking instead of exceptions. Functions that can fail typically return an error as their last return value, which should be checked by the caller. This may seem verbose compared to TypeScript's try/catch, but it makes error paths explicit and helps avoid overlooked errors.

## Building and Running Go Programs

```bash
# Run a Go program directly
go run main.go

# Build an executable
go build main.go

# Install a Go program
go install github.com/user/project

# Test your code
go test ./...

# Run benchmarks
go test -bench=.
```

**Explanation:** Go's tooling is more integrated than TypeScript's ecosystem. The `go` command handles building, testing, formatting, and dependency management. There's no need for separate tools like npm, tsc, jest, and prettier.

## Project Structure

A typical Go project structure:

```
myproject/
├── go.mod              # Module definition
├── go.sum              # Dependencies checksum
├── main.go             # Entry point
├── cmd/                # Command applications
│   └── server/
│       └── main.go
├── internal/           # Private code
│   ├── models/
│   ├── services/
│   └── utils/
├── pkg/                # Public libraries
│   ├── auth/
│   └── database/
└── api/                # API definitions
    └── openapi.yaml
```

**Explanation:** Go projects typically follow a more standardized structure compared to TypeScript/Node.js applications. The `internal` directory contains code that's private to the project, while `pkg` contains code intended to be imported by other projects. This organizational pattern helps clarify the API boundaries.

## Advanced Go Topics for Professional Development

### 1. Advanced Concurrency Patterns

```go
// Worker pools
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}

// Context for cancellation and deadlines
func doWork(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Work cancelled:", ctx.Err())
            return
        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go doWork(ctx)
    time.Sleep(3 * time.Second)
}
```

**Explanation:** Advanced concurrency patterns help solve complex parallel processing problems. Worker pools distribute work across multiple goroutines, while the context package provides a standard way to handle cancellation, deadlines, and request-scoped values. These patterns are essential for building robust, scalable services.

### 2. Reflection

```go
func inspectVariable(v interface{}) {
    value := reflect.ValueOf(v)
    typeOf := reflect.TypeOf(v)

    fmt.Println("Type:", typeOf)

    if typeOf.Kind() == reflect.Struct {
        for i := 0; i < value.NumField(); i++ {
            field := typeOf.Field(i)
            fieldValue := value.Field(i)
            fmt.Printf("Field: %s, Value: %v, Tag: %v\n",
                       field.Name, fieldValue, field.Tag)
        }
    }
}

// Struct tags (similar to decorators in TypeScript)
type User struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"-" validate:"required,min=8"`
}
```

**Explanation:** Reflection in Go allows examining types and values at runtime, similar to TypeScript's type capabilities but at runtime. While not as commonly used as in dynamic languages, reflection is valuable for implementing generic behaviors like ORM libraries, validation, and serialization. Struct tags provide metadata for fields, often used by JSON encoders and validation libraries.

### 3. Testing and Benchmarking

```go
// Basic test (save in *_test.go file)
func TestAdd(t *testing.T) {
    sum := Add(2, 3)
    if sum != 5 {
        t.Errorf("Expected 5, got %d", sum)
    }
}

// Table-driven tests
func TestMultiply(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 6},
        {-1, 5, -5},
        {0, 10, 0},
    }

    for _, test := range tests {
        result := Multiply(test.a, test.b)
        if result != test.expected {
            t.Errorf("Multiply(%d, %d): expected %d, got %d",
                    test.a, test.b, test.expected, result)
        }
    }
}

// Benchmarks
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(20)
    }
}

// Mocks and test helpers
func TestUserService(t *testing.T) {
    // Setup mock database
    mockDB := &MockDatabase{}
    mockDB.On("GetUser", 1).Return(&User{ID: 1, Name: "Test"}, nil)

    service := NewUserService(mockDB)
    user, err := service.GetUser(1)

    assert.NoError(t, err)
    assert.Equal(t, "Test", user.Name)
    mockDB.AssertExpectations(t)
}
```

**Explanation:** Go has a built-in testing framework that's simple yet powerful. Table-driven tests are a common pattern to test multiple inputs and expected outputs. Benchmarking is integrated into the testing package, making it easy to measure performance. For mocking, many Go developers prefer to write simple mock implementations rather than using complex mocking frameworks.

### 4. Code Generation

```go
//go:generate stringer -type=Status
type Status int

const (
    StatusPending Status = iota
    StatusActive
    StatusInactive
)

// Running go generate will create a String() method for Status type
```

**Explanation:** Go's code generation is simpler than TypeScript's decorators or type transformations but is powerful for creating repetitive code. The `go generate` command can run tools to generate code before compilation. Common uses include creating string representations of constants, generating marshaling code for custom types, and creating mock implementations for interfaces.

### 5. Advanced Memory Management

```go
// Memory profiling
import "runtime/pprof"

func main() {
    f, _ := os.Create("mem.prof")
    defer f.Close()

    // Do some work

    pprof.WriteHeapProfile(f)
}

// Escape analysis and stack allocation
func createUserOnStack() User {
    user := User{Name: "John"} // Allocated on stack
    return user
}

func createUserEscapingToHeap() *User {
    user := User{Name: "John"} // Escapes to heap
    return &user
}
```

**Explanation:** Go manages memory for you, but understanding how its garbage collector works helps write more efficient code. Go uses escape analysis to determine whether variables can be allocated on the stack (fast, automatically freed) or must be on the heap (garbage collected). The pprof tools help identify memory leaks and optimize allocation patterns.

### 6. Web Development

```go
// Basic HTTP server
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})
http.ListenAndServe(":8080", nil)

// Using the Gin framework
func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen on port 8080
}

// Middleware
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        c.Next()
        latency := time.Since(t)
        log.Printf("Request took %s", latency)
    }
}
```

**Explanation:** Go excels at web services with its standard library providing robust HTTP support. While frameworks like Gin, Echo, and Fiber offer higher-level abstractions similar to Express in Node.js, many Go developers prefer using the standard library with minimal dependencies. The middleware pattern is common across frameworks for adding cross-cutting concerns like logging, authentication, and error handling.

### 7. Database Access

```go
// SQL with standard library
db, err := sql.Open("postgres", "postgresql://user:pass@localhost/mydb")
if err != nil {
    log.Fatal(err)
}

rows, err := db.Query("SELECT id, name FROM users WHERE active = $1", true)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID: %d, Name: %s\n", id, name)
}

// Using GORM (ORM library)
db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost/mydb"), &gorm.Config{})
if err != nil {
    log.Fatal(err)
}

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"uniqueIndex"`
}

db.AutoMigrate(&User{})
db.Create(&User{Name: "John", Email: "john@example.com"})

var users []User
db.Where("name LIKE ?", "J%").Find(&users)
```

**Explanation:** Database access in Go ranges from low-level SQL with the standard library to higher-level ORMs like GORM. The `database/sql` package provides a consistent interface for different database drivers. Unlike TypeScript ORMs (TypeORM, Prisma), Go ORMs typically require less configuration and have a closer relationship to SQL.

### 8. Performance Optimization

```go
// Avoid allocations in tight loops
var buffer [64]byte // Stack-allocated buffer

// Profile CPU usage
import "runtime/pprof"

func main() {
    f, _ := os.Create("cpu.prof")
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    // Run code to profile
    heavyComputation()
}

// Use sync.Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 4096)
    },
}

func process() {
    buffer := bufferPool.Get().([]byte)
    defer bufferPool.Put(buffer)

    // Use buffer
}
```

**Explanation:** Performance optimization in Go often focuses on reducing allocations and improving memory usage patterns. The Go profiler helps identify bottlenecks, while tools like `sync.Pool` help reuse allocations for better performance. Unlike TypeScript where performance bottlenecks are often in DOM operations or external APIs, Go performance work typically focuses on memory management and algorithm efficiency.

### 9. Cross-Platform Development

```go
// Conditional compilation
// +build linux darwin

package main

// Platform-specific code
func platformSpecificFunction() string {
    return "Running on Linux or macOS"
}

// OS-specific file
// file: file_windows.go
func getPathSeparator() string {
    return "\\"
}

// file: file_unix.go
func getPathSeparator() string {
    return "/"
}
```

**Explanation:** Go excels at cross-platform development with built-in support for conditional compilation based on operating system, architecture, or custom build tags. Unlike TypeScript where platform-specific code often requires runtime checks, Go can include or exclude code at compile time, creating optimized binaries for each target platform.

### 10. Dependency Injection and Application Structure

```go
// Simple dependency injection
type UserRepository interface {
    GetByID(id int) (*User, error)
    Create(user *User) error
}

type UserService struct {
    repo UserRepository
    logger Logger
}

func NewUserService(repo UserRepository, logger Logger) *UserService {
    return &UserService{
        repo: repo,
        logger: logger,
    }
}

// Using wire for dependency injection
// +build wireinject

func InitializeUserHandler(db *sql.DB) (*UserHandler, error) {
    wire.Build(
        NewUserRepository,
        NewUserService,
        NewUserHandler,
    )
    return &UserHandler{}, nil
}
```

**Explanation:** Go applications often use manual dependency injection rather than frameworks like Angular's DI system. This leads to more explicit code but requires more boilerplate. Tools like Google's Wire can generate the necessary wiring code at compile time. The dependency injection pattern helps with testing and maintainability, allowing components to be developed and tested in isolation.

## Your First Go Program

Create a file called `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run hello.go
```

**Explanation:** Every Go program starts with a `package` declaration. The `main` package is special as it defines an executable rather than a library. The `main` function serves as the entry point, similar to TypeScript applications. Unlike TypeScript, there's no need for a separate compilation step during development - `go run` compiles and runs the program in one step.

## Helpful Resources

- [Go by Example](https://gobyexample.com/) - Examples of common Go patterns
- [Tour of Go](https://tour.golang.org/) - Interactive tutorial
- [Effective Go](https://golang.org/doc/effective_go) - Official best practices
- [Go Playground](https://play.golang.org/) - Test Go code in browser
- [Go Modules Reference](https://golang.org/ref/mod) - Guide to dependency management
- [Standard Library Documentation](https://golang.org/pkg/) - Comprehensive standard library docs
- [The Go Blog](https://blog.golang.org/) - Official Go blog with in-depth articles
