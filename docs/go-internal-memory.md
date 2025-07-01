Go **internal memory layout**—specifically the **Code Segment, Data Segment, Stack, and Heap**—and how each is used by the Go runtime.

---

## 📦 1. Code Segment

- **What it is**: Stores the compiled machine code (functions, methods, control flow).
- **Characteristics**:

  - Read-only during runtime.
  - Loaded into memory when the Go program starts.

- **Example contents**:

  - The actual instructions of `main()` function.
  - Runtime code (e.g., garbage collector, scheduler).

- **Managed by**: OS and Go runtime during program startup.

---

## 📊 2. Data Segment

The data segment contains **global and static variables**. It is typically split into two parts:

### 🔹 a. **Initialized Data**

- Stores global variables that have an explicit initial value.
- Allocated at program load time and initialized with those values.

### 🔹 b. **Uninitialized Data (BSS)**

- Stores global variables declared but not explicitly initialized (e.g., `var x int`).
- Initialized to zero by default.

> 🧠 These variables live for the entire program duration.

---

## 🧮 3. Stack

- **What it is**: Stores function-local variables, return addresses, and call stack frames.
- **Characteristics**:

  - Allocate **Stack Frame** for each function.
  - Each goroutine has its **own stack** (starts small, grows/shrinks dynamically).
  - Very fast allocation/deallocation (LIFO).

- **Managed by**: Go runtime and scheduler.

> Variables on the stack are **not managed by the garbage collector** (unless they escape to the heap).

**Example**:

```go
func foo() {
    x := 42   // allocated on the stack (if it doesn’t escape)
}
```

---

## 🗃️ 4. Heap

- **What it is**: Stores dynamically allocated memory and long-lived variables.
- **Characteristics**:

  - Shared across all goroutines.
  - Managed by the **garbage collector**.
  - Slower allocation than stack, but necessary for persistent data.

**Example**:

```go
func bar() *int {
    x := 42
    return &x // x escapes to the heap
}
```

---

## Visual Overview (Simplified Memory Layout)

```
|-----------------------------|
|         Code Segment        |  ← machine instructions
|-----------------------------|
|         Data Segment        |  ← global/static variables
|-----------------------------|
|            Heap             |  ← dynamic memory (GC managed)
|-----------------------------|
|            Stack            |  ← local variables, call frames (per goroutine)
|-----------------------------|
```

---

## Summary Table

| Segment | Stores                       | Lifetime       | Managed by                 |
| ------- | ---------------------------- | -------------- | -------------------------- |
| Code    | Instructions                 | Entire program | OS / Go runtime            |
| Data    | Global/static vars           | Entire program | Go runtime                 |
| Stack   | Local vars, return addresses | Function call  | Go runtime (per goroutine) |
| Heap    | Dynamic memory, escaped vars | Varies         | Garbage Collector          |

---

Let me know if you'd like a diagram, example with memory addresses, or how to observe this using Go tooling!
