# Advanced Concurrency Patterns in Go

## Introduction to Go Concurrency

Concurrency is one of Go's most powerful features and a major reason many developers choose Go for building high-performance systems. Unlike JavaScript/TypeScript where asynchronous programming is built around callbacks, promises, and async/await, Go provides native language constructs specifically designed for concurrent programming.

The foundation of Go's concurrency model is built on three key concepts:
1. **Goroutines** - Lightweight threads managed by the Go runtime
2. **Channels** - Typed conduits for sending and receiving values between goroutines
3. **Select** - A control structure for working with multiple channel operations

For someone new to Go concurrency, think of goroutines as independent tasks that can run concurrently, and channels as pipes that connect these tasks, allowing them to communicate safely.

## 1. Worker Pools

Worker pools are a common concurrency pattern where a fixed number of worker goroutines process jobs from a shared channel. This pattern is ideal for CPU-intensive tasks or when you need to limit concurrent operations (like database connections or API calls).

### Basic Worker Pool Implementation

```go
package main

import (
	"fmt"
	"time"
)

// Task represents work to be done
type Task struct {
	ID      int
	JobData string
	Result  string
}

func worker(id int, tasksChan <-chan Task, resultsChan chan<- Task) {
	fmt.Printf("Worker %d starting\n", id)
	
	// Process tasks until the channel is closed
	for task := range tasksChan {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		
		// Simulate work
		time.Sleep(500 * time.Millisecond)
		
		// Update the task with results
		task.Result = fmt.Sprintf("Processed by worker %d", id)
		
		// Send the result back
		resultsChan <- task
	}
	
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create channels for tasks and results
	tasksChan := make(chan Task, 10) // Buffered channel with capacity 10
	resultsChan := make(chan Task, 10)
	
	// Start 3 workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasksChan, resultsChan)
	}
	
	// Send tasks
	numTasks := 5
	for i := 1; i <= numTasks; i++ {
		tasksChan <- Task{ID: i, JobData: fmt.Sprintf("Task %d data", i)}
	}
	
	// Close the tasks channel - this tells the workers that no more tasks are coming
	close(tasksChan)
	
	// Collect results
	for i := 1; i <= numTasks; i++ {
		result := <-resultsChan
		fmt.Printf("Result: Task %d was %s\n", result.ID, result.Result)
	}
}
```

### Key Concepts in Worker Pools:

1. **Channel Directionality**:
   ```go
   tasksChan <-chan Task   // Read-only channel (worker can only receive)
   resultsChan chan<- Task // Write-only channel (worker can only send)
   ```
   This makes the intent clear and prevents mistakes like closing a channel from the wrong side.

2. **Buffered Channels**:
   ```go
   tasksChan := make(chan Task, 10)
   ```
   Allows sending up to 10 values without blocking, which can improve performance by reducing goroutine switching.

3. **Worker Management**: Start a fixed number of workers based on available resources:
   ```go
   numWorkers := runtime.NumCPU() // Use number of CPUs as a guide
   ```

4. **Signaling Completion**: Close the input channel to signal that no more tasks are coming:
   ```go
   close(tasksChan)
   ```

### Real-world Example: Processing Images

```go
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"
)

type ImageTask struct {
	InputPath  string
	OutputPath string
	Error      error
}

// Process an image - resize or apply filters
func processImage(task ImageTask) ImageTask {
	// Open the image
	file, err := os.Open(task.InputPath)
	if err != nil {
		task.Error = fmt.Errorf("failed to open image: %w", err)
		return task
	}
	defer file.Close()
	
	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		task.Error = fmt.Errorf("failed to decode image: %w", err)
		return task
	}
	
	// Create the output directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(task.OutputPath), os.ModePerm)
	if err != nil {
		task.Error = fmt.Errorf("failed to create output directory: %w", err)
		return task
	}
	
	// Create the output file
	outFile, err := os.Create(task.OutputPath)
	if err != nil {
		task.Error = fmt.Errorf("failed to create output file: %w", err)
		return task
	}
	defer outFile.Close()
	
	// In a real application, you would transform the image here
	// For example: resize, apply filters, etc.
	
	// Write the image to the output file
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		task.Error = fmt.Errorf("failed to encode image: %w", err)
		return task
	}
	
	return task
}

func worker(id int, tasks <-chan ImageTask, results chan<- ImageTask) {
	for task := range tasks {
		fmt.Printf("Worker %d processing %s\n", id, task.InputPath)
		results <- processImage(task)
	}
}

func main() {
	// Create input/output paths (simplified for the example)
	inputDir := "./input"
	outputDir := "./output"
	
	// Create channels
	tasks := make(chan ImageTask, 100)
	results := make(chan ImageTask, 100)
	
	// Start workers
	numWorkers := 4
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			worker(id, tasks, results)
			wg.Done()
		}(i)
	}
	
	// Create a goroutine to close the results channel once all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Send tasks (in a real app, you would scan the input directory)
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- ImageTask{
				InputPath:  filepath.Join(inputDir, fmt.Sprintf("image%d.jpg", i)),
				OutputPath: filepath.Join(outputDir, fmt.Sprintf("processed%d.jpg", i)),
			}
		}
		close(tasks)
	}()
	
	// Collect results
	successCount := 0
	failCount := 0
	
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error processing %s: %v\n", result.InputPath, result.Error)
			failCount++
		} else {
			fmt.Printf("Successfully processed %s -> %s\n", result.InputPath, result.OutputPath)
			successCount++
		}
	}
	
	fmt.Printf("Processing complete. Success: %d, Failed: %d\n", successCount, failCount)
}
```

This example demonstrates a more complete worker pool implementation for image processing, including error handling and proper coordination of goroutines.

## 2. Context Management

The `context` package provides a way to carry deadlines, cancellation signals, and request-scoped values across API boundaries and between processes. It's crucial for building robust, production-ready applications.

### Basic Context Usage

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) (string, error) {
	// Create a channel for our result
	resultCh := make(chan string)
	
	// Start the actual work in a separate goroutine
	go func() {
		// Simulate work that takes time
		time.Sleep(2 * time.Second)
		resultCh <- "Operation completed successfully"
	}()
	
	// Wait for either the operation to complete or the context to be cancelled
	select {
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err() // This will be either Canceled or DeadlineExceeded
	}
}

func main() {
	// Create a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Always call cancel to release resources
	
	fmt.Println("Starting operation...")
	result, err := longRunningOperation(ctx)
	
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}
```

This program creates a context with a 1-second timeout and passes it to a function that simulates work taking 2 seconds. Since the operation takes longer than the timeout, the context will be cancelled, and the function will return an error.

### Types of Contexts

1. **Background Context**:
   ```go
   ctx := context.Background()
   ```
   This is the root context, typically used at the start of a program or request.

2. **TODO Context**:
   ```go
   ctx := context.TODO()
   ```
   Used when it's unclear which context to use or when the function will be updated to accept a context.

3. **WithCancel Context**:
   ```go
   ctx, cancel := context.WithCancel(parentCtx)
   // Later, call cancel() to stop operations
   ```
   Allows manual cancellation of operations.

4. **WithTimeout Context**:
   ```go
   ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
   ```
   Automatically cancels after the specified duration.

5. **WithDeadline Context**:
   ```go
   ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Minute))
   ```
   Cancels at a specific time.

6. **WithValue Context**:
   ```go
   ctx := context.WithValue(parentCtx, key, value)
   ```
   Carries request-scoped values across API boundaries.

### Real-world Example: HTTP Server with Context

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// UserService simulates a database service
type UserService struct {
	// In a real app, this would be a database connection
}

// GetUser fetches a user by ID using context for cancellation
func (s *UserService) GetUser(ctx context.Context, id string) (map[string]interface{}, error) {
	// Create a channel for our result
	resultCh := make(chan map[string]interface{})
	errCh := make(chan error)
	
	go func() {
		// Simulate a database query
		time.Sleep(500 * time.Millisecond)
		
		// Check if the context was cancelled before responding
		select {
		case <-ctx.Done():
			// Context was cancelled, don't bother sending a result
			return
		default:
			// Context is still valid, send the result
			resultCh <- map[string]interface{}{
				"id":    id,
				"name":  "John Doe",
				"email": "john@example.com",
			}
		}
	}()
	
	// Wait for either the operation to complete or the context to be cancelled
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	userService := &UserService{}
	
	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		// Extract user ID from URL
		userID := r.URL.Path[len("/api/users/"):]
		
		// Create a context that will be cancelled when the client disconnects
		ctx := r.Context()
		
		// Set a timeout for the database operation
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		
		// Get user data
		user, err := userService.GetUser(ctx, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Return the result as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	
	// Start the server
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

This example shows how contexts are used in real-world HTTP servers to manage request timeouts and cancellations. HTTP requests in Go already have an associated context that is cancelled when the client disconnects.

## 3. Synchronization Primitives

Go provides several synchronization primitives in the `sync` package for coordinating goroutines. These are useful when channels are not the most appropriate solution.

### Mutex: Protecting Shared Resources

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	
	// Start 1000 goroutines that increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	// Wait for all goroutines to finish
	wg.Wait()
	
	fmt.Printf("Final counter value: %d\n", counter.Value())
}
```

Without the mutex, the counter would likely end up with a value less than 1000 due to race conditions.

### RWMutex: Read-Write Locking

When many goroutines read from a shared resource and few write to it, `RWMutex` can improve performance by allowing multiple readers concurrently.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		items: make(map[string]string),
	}
}

func (m *SafeMap) Get(key string) (string, bool) {
	m.mu.RLock() // Multiple readers can acquire the lock
	defer m.mu.RUnlock()
	val, ok := m.items[key]
	return val, ok
}

func (m *SafeMap) Set(key, value string) {
	m.mu.Lock() // Only one writer can acquire the lock
	defer m.mu.Unlock()
	m.items[key] = value
}

func main() {
	cache := NewSafeMap()
	var wg sync.WaitGroup
	
	// Simulate many readers and few writers
	
	// Start 10 writer goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			cache.Set(key, value)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	// Start 100 reader goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key-%d", j%10)
				val, _ := cache.Get(key)
				_ = val // Use val to avoid compiler warning
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	
	wg.Wait()
	fmt.Println("All operations completed")
}
```

### WaitGroup: Waiting for Goroutines

`WaitGroup` is used to wait for a collection of goroutines to finish. We've already seen it in action in previous examples.

```go
var wg sync.WaitGroup

wg.Add(1)  // Add 1 to the counter
go func() {
    defer wg.Done()  // Decrement counter when goroutine completes
    // Do work
}()

wg.Wait()  // Block until counter becomes 0
```

### Once: Run Something Only Once

`sync.Once` ensures that a function is executed only once, even when called from multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	
	// This initialization function will only be called once
	initialize := func() {
		fmt.Println("Initializing...")
	}
	
	// Start 10 goroutines that all try to initialize
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d trying to initialize\n", id)
			once.Do(initialize)
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All goroutines finished")
}
```

This is useful for lazy initialization of resources like database connections.

### Cond: Condition Variables

`sync.Cond` provides a way for goroutines to wait for a condition to become true.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a condition variable with an associated mutex
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	
	// Create a queue
	queue := make([]int, 0, 10)
	
	// Start consumer goroutines
	for i := 0; i < 2; i++ {
		go func(id int) {
			for {
				// Acquire the lock
				mu.Lock()
				
				// Wait until there's an item in the queue
				for len(queue) == 0 {
					fmt.Printf("Consumer %d: waiting for items\n", id)
					cond.Wait() // Releases lock and waits; reacquires lock when woken
				}
				
				// Get an item from the queue
				item := queue[0]
				queue = queue[1:]
				
				// Release the lock
				mu.Unlock()
				
				// Process the item
				fmt.Printf("Consumer %d: processed item %d\n", id, item)
				time.Sleep(time.Second)
			}
		}(i)
	}
	
	// Produce items
	for i := 1; i <= 5; i++ {
		// Acquire the lock
		mu.Lock()
		
		// Add an item to the queue
		queue = append(queue, i)
		fmt.Printf("Producer: added item %d to queue\n", i)
		
		// Signal one goroutine that's waiting
		cond.Signal()
		
		// Release the lock
		mu.Unlock()
		
		time.Sleep(time.Second)
	}
	
	// Add a few more items all at once
	mu.Lock()
	for i := 6; i <= 8; i++ {
		queue = append(queue, i)
		fmt.Printf("Producer: added item %d to queue\n", i)
	}
	
	// Wake up all waiting goroutines
	cond.Broadcast()
	mu.Unlock()
	
	// Let the program run for a while
	time.Sleep(10 * time.Second)
}
```

`Cond.Wait()` atomically releases the associated mutex and suspends the goroutine. When `Cond.Signal()` or `Cond.Broadcast()` is called, the goroutine is woken up, reacquires the mutex, and continues execution.

## 4. Fan-out, Fan-in Pattern

This pattern involves "fanning out" work to multiple goroutines and then "fanning in" the results.

```go
package main

import (
	"fmt"
	"sync"
)

// generator creates a channel that emits numbers from 1 to n
func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()
	return out
}

// square calculates the square of numbers received from the in channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// merge combines multiple channels into a single channel
func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	
	// Start an output goroutine for each input channel
	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			out <- n
		}
	}
	
	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}
	
	// Start a goroutine to close the output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func main() {
	// Create a channel of numbers
	nums := generator(10)
	
	// Fan out: distribute work across multiple goroutines
	sq1 := square(nums)
	sq2 := square(nums)
	sq3 := square(nums)
	
	// Fan in: combine results
	for n := range merge(sq1, sq2, sq3) {
		fmt.Println(n)
	}
}
```

Wait, there's a problem with the above example! Since channels in Go can't be cloned, only one goroutine will receive each number from the generator. Let's fix that:

```go
package main

import (
	"fmt"
	"sync"
)

// generator creates a channel that emits numbers from 1 to n
func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()
	return out
}

// square calculates the square of numbers received from the in channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// distribute fans out values from in to multiple output channels
func distribute(in <-chan int, n int) []<-chan int {
	outputs := make([]<-chan int, n)
	
	// Create n channels
	for i := 0; i < n; i++ {
		outputs[i] = make(chan int)
	}
	
	// Start a goroutine to distribute values
	go func() {
		defer func() {
			// Close all output channels when done
			for i := 0; i < n; i++ {
				close(outputs[i].(chan int))
			}
		}()
		
		i := 0
		for val := range in {
			// Send the value to the next channel in a round-robin fashion
			outputs[i].(chan int) <- val
			i = (i + 1) % n
		}
	}()
	
	return outputs
}

// merge combines multiple channels into a single channel
func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	
	// Start an output goroutine for each input channel
	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			out <- n
		}
	}
	
	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}
	
	// Start a goroutine to close the output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func main() {
	// Create a channel of numbers
	nums := generator(10)
	
	// Fan out: distribute work across multiple goroutines
	numChannels := 3
	distributedNums := distribute(nums, numChannels)
	
	squares := make([]<-chan int, numChannels)
	for i := 0; i < numChannels; i++ {
		squares[i] = square(distributedNums[i])
	}
	
	// Fan in: combine results
	for n := range merge(squares...) {
		fmt.Println(n)
	}
}
```

Hmm, there's still an issue with type assertions in the `distribute` function. Let's fix it properly:

```go
package main

import (
	"fmt"
	"sync"
)

// generator creates a channel that emits numbers from 1 to n
func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()
	return out
}

// square calculates the square of numbers received from the in channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// distribute fans out values from in to multiple output channels
func distribute(in <-chan int, n int) []<-chan int {
	outputs := make([]chan int, n)
	for i := 0; i < n; i++ {
		outputs[i] = make(chan int)
	}
	
	// Convert to read-only channels for return
	readOnlyOutputs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		readOnlyOutputs[i] = outputs[i]
	}
	
	// Start a goroutine to distribute values
	go func() {
		defer func() {
			// Close all output channels when done
			for i := 0; i < n; i++ {
				close(outputs[i])
			}
		}()
		
		i := 0
		for val := range in {
			// Send the value to the next channel in a round-robin fashion
			outputs[i] <- val
			i = (i + 1) % n
		}
	}()
	
	return readOnlyOutputs
}

// merge combines multiple channels into a single channel
func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	
	// Start an output goroutine for each input channel
	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			out <- n
		}
	}
	
	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}
	
	// Start a goroutine to close the output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func main() {
	// Create a channel of numbers
	nums := generator(10)
	
	// Fan out: distribute work across multiple goroutines
	numWorkers := 3
	distributedNums := distribute(nums, numWorkers)
	
	squares := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		squares[i] = square(distributedNums[i])
	}
	
	// Fan in: combine results
	for n := range merge(squares...) {
		fmt.Println(n)
	}
}
```

This pattern is useful for CPU-bound tasks where you want to distribute work across multiple cores.

## 5. Rate Limiting

Rate limiting is essential for controlling the rate of requests to a resource. It can be implemented elegantly using channels as a token bucket.

```go
package main

import (
	"fmt"
	"time"
)

// RateLimiter implements a token bucket rate limiter
type RateLimiter struct {
	tokens  chan struct{}
	ticker  *time.Ticker
	stop    chan struct{}
	stopped bool
}

// NewRateLimiter creates a new rate limiter with the given rate (tokens per second)
// and burst size (maximum tokens at once)
func NewRateLimiter(rate float64, burst int) *RateLimiter {
	// Calculate interval between tokens
	interval := time.Duration(float64(time.Second) / rate)
	
	limiter := &RateLimiter{
		tokens: make(chan struct{}, burst),
		ticker: time.NewTicker(interval),
		stop:   make(chan struct{}),
	}
	
	// Add initial tokens
	for i := 0; i < burst; i++ {
		limiter.tokens <- struct{}{}
	}
	
	// Refill tokens at the specified rate
	go func() {
		for {
			select {
			case <-limiter.ticker.C:
				// Try to add a token without blocking
				select {
				case limiter.tokens <- struct{}{}:
					// Token added
				default:
					// Bucket full, discard token
				}
			case <-limiter.stop:
				limiter.ticker.Stop()
				return
			}
		}
	}()
	
	return limiter
}

// Wait blocks until a token is available
func (r *RateLimiter) Wait() {
	<-r.tokens
}

// Stop stops the rate limiter
func (r *RateLimiter) Stop() {
	if !r.stopped {
		close(r.stop)
		r.stopped = true
	}
}

func main() {
	// Create a rate limiter with 2 requests per second and burst of 5
	limiter := NewRateLimiter(2, 5)
	defer limiter.Stop()
	
	// Simulate 20 requests
	for i := 1; i <= 20; i++ {
		go func(id int) {
			fmt.Printf("Request %d: waiting for token\n", id)
			
			start := time.Now()
            limiter.Wait()
            elapsed := time.Since(start)
            
            fmt.Printf("Request %d: got token after %v\n", id, elapsed)
            
            // Simulate processing the request
            time.Sleep(100 * time.Millisecond)
            
            fmt.Printf("Request %d: completed\n", id)
        }(i)
    }
    
    // Wait to observe the rate limiting in action
    time.Sleep(12 * time.Second)
}
```

This Go implementation demonstrates a token bucket rate limiter using Go's concurrency primitives:

1. **Token Bucket Algorithm**: The rate limiter uses a buffered channel (`tokens`) as a token bucket. The bucket has a maximum capacity (burst size) and is refilled at a steady rate.

2. **Implementation Details**:
   - The `NewRateLimiter` function creates a limiter with the specified rate and burst size
   - Tokens are implemented as empty structs (`struct{}{}`) which use minimal memory
   - A goroutine continuously adds tokens at the specified rate
   - If the bucket is full, additional tokens are discarded

3. **Usage Pattern**:
   - The `Wait()` method blocks until a token is available
   - When a request needs to be processed, it calls `Wait()` to acquire a token
   - If no tokens are available, the request waits until a token is added

4. **Main Function Demo**:
   - Creates a rate limiter allowing 2 requests per second with a burst of 5
   - Simulates 20 concurrent requests
   - Each request waits for a token, processes for 100ms, then completes
   - The program runs for 12 seconds to observe the rate limiting behavior

This implementation elegantly leverages Go's channels for both the token bucket and the synchronization mechanism, making it thread-safe and efficient.
