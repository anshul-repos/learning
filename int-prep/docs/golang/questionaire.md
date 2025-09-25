# Golang Notes

## 1. GO Latest Version
- **1.25**

---

## 2. Go Modules and Packages
- A **module** is a collection of Go packages.
- A **package** is a directory of `.go` files.

---

## 3. Fallthrough Example
```go
func main() {
    day := "Tuesday"
    switch day {
    case "Monday":
        fmt.Println("It's Monday.")
        fallthrough
    case "Tuesday":
        fmt.Println("It's Tuesday.")
        fallthrough
    case "Wednesday":
        fmt.Println("It's Wednesday.")
    default:
        fmt.Println("It's some other day.")
    }
}
```
**Output:**
```
It's Tuesday.
It's Wednesday.
It's some other day.
```

---

## 4. Slices
- Similar to arrays, but more powerful and flexible.
- **Slice Structure:** Length, capacity, and pointer.
- **Slice Length:** Number of elements in the slice.
- **Slice Capacity:** Maximum elements the slice can grow without reallocating.

---

## 5. Embedded Struct
- Go allows structs to be embedded within other structs (composition).

---

## 6. Binary File
- A binary file is an executable file generated after you build your Go source code.
- Platform-specific, contains machine code executable directly by the OS.

---

## 7. Goroutine vs Thread
| Goroutine | Thread |
|-----------|--------|
| 2KB size | 1MB size |
| Managed by Go runtime | Managed by kernel |
| Not hardware dependent | Hardware dependent |
| Easy communication via channels | No easy communication medium |
| Low latency inter-goroutine communication | High latency inter-thread communication |
| No ID (Go lacks TLS) | Unique ID (has TLS) |
| Cheaper than threads | More costly |
| Cooperatively scheduled | Preemptively scheduled |
| Faster startup time | Slower startup time |
| Fast context switching | Slow context switching |

---

## 8. Functions
- First-class functions
- Higher-order functions
- Closures
- Anonymous functions
- Variadic functions

---

### 8.1 Variadic Functions
- Functions that accept a variable number of arguments using `...`
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

### 8.2 Closures

- A special kind of anonymous function that captures variables from its surrounding scope and can use them even after that scope has ended.
- The captured variables live as long as the closure itself lives.

```go
func main() {
    counter := func() func(int) int {
        count := 0
        return func(x int) int {
            count += x
            return count
        }
    }
    c := counter()
    fmt.Println(c(1)) // 1
    fmt.Println(c(2)) // 3
}
```
Here:

- The returned function is anonymous, but it also captures and remembers count from counter().
- Even after counter() has finished executing, the returned function still has access to count.
- This is exactly what a closure does: it “closes over” variables from its surrounding scope.

---

### 8.3 Anonymous Function

- A function without a name.
- It can be assigned to a variable, passed as an argument, or used immediately (IIFE — Immediately Invoked Function Expression).
- By itself, it doesn’t necessarily capture variables from outside its scope.
- Example:
```go
f := func(a, b int) int {
    return a + b
}
fmt.Println(f(2, 3)) // 5
```
- Here, func(a, b int) int { ... } is anonymous because it has no name.

---

## 9. Interfaces in Go
- An interface is a set of method signatures.
- A value of interface type can hold any value implementing those methods.

---

## 10. Variable Initialization
- `var result []string` → does not allocate memory.
- `result := []string{}` → allocates memory for an empty slice.

---

## 11. Array vs Slice
| Array | Slice |
|-------|-------|
| `var arr [10]int` | `var slice []int` |
| Fixed size | Dynamic size |
| Has length only | Has length & capacity |
| Cannot use `make()` | Can use `make()` |
| Values must be initialized | Optional initialization |
| Cannot be resized | Resizable via `append()` |

---

## 12. Go Data Types
- **Fundamental:** int, float, bool, string, rune
- **Composite:** array, slice, struct, map, channel
- **Reference:** pointer, function, interface

---

## 13. make() vs new()
- **make:**
- Allocates and initializes memory for `slices, maps, channels only` (cannot be used for arrays or structs).
- Returns value.
    ```go
    s := make([]int, 3)
    ```
- **new:** 
- Allocates memory, does not initialize.
- For simple types like `int, struct, float64, etc`
- Returns pointer.
  ```go
  p := new(int)
  ````
  -  p is of type *int
---

## 14. Delete a Key from Map
```go
delete(mapName, keyName)
delete(numMap, num)
```

---

## 15. Functions vs Methods
| Feature | Function | Method |
|----------|----------|--------|
| Receiver | No receiver | Requires receiver |
| Binding | Independent | Bound to type |
| Purpose | General tasks | Defines behavior for type |
| Call Syntax | `Add(2,3)` | `rect.Area()` |

---

## 16. strings.Fields vs strings.Split
- `strings.Fields`: Splits by whitespace.
- `strings.Split`: Splits by a delimiter.

---

## 17. Slice Pointer Types
| Type | Meaning | Use Case |
|------|---------|----------|
| `*[]T` | Pointer to slice | Modify slice itself |
| `[]*T` | Slice of pointers | Modify elements via pointer |

- `Use *[]T` when you want to change the slice container itself (grow, reassign).

- `Use []*T` when you want to change the values that the slice points to.

---

## 18. Check if Map Contains Key
```go
value, ok := myMap[key]
```

---

## 19. Shadowing in Go
- Shadowing happens when you declare a new variable with the same name as an existing variable in an inner scope.
- The new variable “shadows” the outer one, making the outer variable inaccessible in that scope.
- Example:
```go
x := 10
fmt.Println("Outer x:", x)
{
    x := 20 // shadows outer x
    fmt.Println("Inner x:", x)
}
fmt.Println("After block, Outer x:", x)
```

Output:
```
Outer x: 10
Inner x: 20
After block, Outer x: 10
```

---

## 20. Synchronous vs Asynchronous
- **Synchronous:** Blocking, sequential.
- **Asynchronous:** Non-blocking, often with goroutines.

---

## 21. Performance Optimization in Go
- Use `int` instead of `string` in maps.
- Sort struct fields (use `fieldalignment`).
- Use goroutines.
- Profile with `pprof`.
- Prefer protocol buffers over JSON.
- Pass structs by pointer.
- Preallocate slices.
- Reuse objects instead of reallocating.

---

## 22. String Operations
```go
str := strings.ToLower(s) // lowercase
strconv.Itoa(123)         // int to string
```

- A byte holds ASCII.
- Strings may contain Unicode (loop yields runes).

---

## 23. Garbage Collector (GC)
- Uses **mark and sweep** with tri-color marking.
- Runs concurrently with application.
- Optimized for short-lived and long-lived objects.

---

## 24. Context
- Carries request-scoped data, deadlines, cancellation signals.
- Types:
  - `context.Background()`
  - `context.TODO()`
  - `context.WithCancel()`
  - `context.WithTimeout()`
  - `context.WithDeadline()`
  - `context.WithValue()`

---

## 25. Channels
```go
c := make(chan string) // unbuffered
c := make(chan string, 2) // buffered
```
- Send: `c <- "ping"`
- Receive: `msg := <-c`
- Channel Direction:
  - `chan<-` (send only)
  - `<-chan` (receive only)

---

## 26. Pointer to Pointer

- `var v int = 100` → A normal integer.
- `var p1 *int = &v` → p1 is a pointer to v.
- `var p2 **int = &p1` → p2 is a pointer to a pointer. It points to p1, which in turn points to v. So the chain looks like: `p2 → p1 → v (100)`

```go
var v int = 100
var p1 *int = &v
var p2 **int = &p1
fmt.Println(**p2) // 100
```

---

## 27. Defer, Panic & Recover
- **Defer:** Executes after surrounding function returns.
- **Panic:** Stops execution, triggers defers.
- **Recover:** Recovers from panic inside deferred function.

---

## 28. Nested Goroutines Behavior
- Goroutines inside other goroutines still run concurrently.
- The parent goroutine does **not** wait for the child automatically.
- Use `sync.WaitGroup` or channels to coordinate execution.

---

## 29. Connection Pooling
- Reduces overhead of creating new DB connections.
- In Go (`database/sql`):
  - `db.SetMaxOpenConns(n)` → max open connections.
  - `db.SetMaxIdleConns(n)` → idle connections.
  - `db.SetConnMaxLifetime(d)` → connection lifetime.

---

## 30. Identifying Issues in Long SQL Queries
- Use `EXPLAIN` to analyze query execution plan.
- Look for:
  - Full table scans.
  - Missing indexes.
  - Large temporary tables.
- Profiling tools help identify bottlenecks.

---

## 31. Pagination with LIMIT and OFFSET
```sql
SELECT * FROM users ORDER BY id LIMIT 10 OFFSET 20;
```
- `LIMIT` → number of rows.
- `OFFSET` → starting row.
- Alternative: keyset pagination for better performance on large datasets.

---

## 32. Buffered vs Unbuffered Channels
- **Unbuffered**:
  - No capacity (capacity = 0).
  - Send and receive operations block until the other side is ready.
  - Good for synchronization between goroutines.
  - Example: `ch := make(chan string)` // unbuffered
- **Buffered**: 
  - Have capacity > 0.
  - Sender does not block until buffer is full.
  - A sender can place up to capacity values without blocking.
  - A receiver can take values later, until the buffer is empty.
  - Good for queueing work or decoupling senders and receivers.
  - Example: `ch := make(chan int,10)` // buffered with capacity 2

---

## 33. Deep Copy vs Shallow Copy
- **Shallow copy**:
    - A shallow copy just copies references/pointers to underlying data.
    - Modifying the copied object may also modify the original, since both point to the same memory.
- **Deep copy**: 
    - A deep copy creates a completely new object with its own memory.
    - Modifying the copy does not affect the original.
    - Use deep copy for independent objects.

---

## 34. Garbage Collection in Go
- Uses **tri-color tracing** and **mark-and-sweep**.
- Automatically frees memory of unreachable objects.
- Optimized for low pause times.

---

## 35. Purpose of Go Modules
- Dependency management system.
- Each project has a `go.mod` file.
- Handles versioning and reproducible builds.

---

## 36. Context
- Used for deadlines, cancellations, request-scoped values.
- Functions: `context.WithCancel`, `context.WithDeadline`, `context.WithTimeout`.
- Propagates cancellation across goroutines.

---

## 37. Type Assertion
```go
var i interface{} = "hello"
s := i.(string) // type assertion
```
- Retrieves underlying type from an interface.
- Second value (`ok`) avoids panic.

---

## 38. Reflection Package
- Inspect and manipulate types at runtime.
- `reflect.TypeOf()` and `reflect.ValueOf()`.
- Useful for generic libraries and serialization.

---

## 39. Type Conversion
```go
var x int = 10
var y float64 = float64(x)
```
- Use `T(value)` syntax.
- String to int: `strconv.Atoi("123")`.

---

## 40. Worker Pool / Task Jobs Implementation
- Pattern for limiting concurrent goroutines.
- Example: Workers read jobs from a channel, process, and write results.

---

## 41. Designing High Concurrency System
- Use goroutines + channels.
- Connection pooling.
- Load balancing.
- Caching (Redis, Memcached).
- Avoid global locks.

---

## 42. Coding Example
**Input:** `s = "helloArray"`
**Output:** `h1e1ll2o1A1rr2a1y1`
- Use map to count character frequencies.

---

## 43. Datatypes in Go
- Basic: `int`, `float`, `string`, `bool`.
- Composite: `array`, `slice`, `map`, `struct`.
- Reference: `pointer`, `interface`, `channel`.

---

## 44. Methods vs Functions
- **Function**: Independent, not tied to a type.
- **Method**: Has a receiver, associated with a type.

---

## 45. Features of Go and Concurrency
- Lightweight goroutines.
- Channels for communication.
- Select statement.
- Memory-safe concurrency.

---

## 46. gRPC vs HTTP
- **HTTP/REST**: Human-readable, JSON.
- **gRPC**: Uses Protobuf, faster, binary, HTTP/2.
- gRPC better for microservices.

---

## 47. Structs
- Custom data type grouping fields.
```go
type Person struct {
  Name string
  Age  int
}
```

---

## 48. Exception Handling in Go
- Go uses **errors**, **panic**, **recover**, **defer**.
- Example:
```go
defer func() {
  if r := recover(); r != nil {
    fmt.Println("Recovered:", r)
  }
}()
panic("Something went wrong")
```

---

## 49. String Literals
- **Interpreted** (`"Hello"`) → supports escape chars.
- **Raw** (`` `Hello` ``) → multi-line, no escaping.

---

## 50. `init` in Go
- Runs before `main()`.
- Can have multiple `init()` per package.
- Used for setup.

---

## 51. sync.Map
- Concurrent-safe map.
- Optimized for frequent reads & occasional writes.
- Methods: `Store`, `Load`, `Delete`, `Range`.

---

## 52. Mutex vs Semaphore
- **Mutex**: One goroutine at a time.
- **Semaphore**: Allows multiple goroutines up to a limit.

---

## 53. Deferred Calls
- Executed in **reverse order (LIFO)**.

---

## 54. Goroutines vs Threads
- Goroutines: lightweight, small stack, managed by Go runtime.
- Threads: OS-managed, heavier.

---

## 55. Managing Concurrency & Race Conditions
- Use `sync.Mutex`.
- Detect with `go run -race`.

---

## 56. REST vs gRPC
- REST: Flexible, human-readable.
- gRPC: High performance, strict contracts, binary.

---

## 57. Unit Testing with Ginkgo & Gomega
- Ginkgo → BDD framework.
- Gomega → expressive matchers.
- Mock dependencies, cover edge cases.

---

## 58. Structuring Large Projects
- Follow **Clean Architecture**.
- Folder structure: `/cmd`, `/pkg`, `/internal`, `/api`.
- Use Go modules for versioning.

---

## 59. Common Slice Operations
- Create, index, slice.
- Append elements / slices.
- Copy slices.
- Delete elements.

---

## 60. Panic and Defer
- **Panic**: Stops normal execution.
- **Defer**: Runs before function exit (LIFO).

---

## 61. Error Handling
- Standard: return `error`.
- Custom errors implement `Error()`.
- Example:
```go
type MyError struct {
  Code int
  Msg  string
}
func (e *MyError) Error() string {
  return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
```

---

## 62. Race Conditions & Deadlocks
- **Race**: Multiple goroutines modify shared memory.
- **Deadlock**: Goroutines wait indefinitely.
- Detect race with `-race` flag.

---

## 63. Types of Mutex
- **sync.Mutex**: Exclusive access.
- **sync.RWMutex**: Multiple readers or one writer.

| Feature | sync.Mutex | sync.RWMutex |
|---------|------------|--------------|
| Readers | No         | Yes          |
| Writers | Yes        | Yes          |
| Perf    | Lower      | Higher (read-heavy) |

---

## 64. Common Go (Golang) Pitfalls & How to Avoid Them

This document lists the **most common pitfalls in Go** and how to avoid them.  
It’s useful for developers working on backend services, microservices, or distributed systems.

---

### 64.1 Ignoring Errors
Failing to check and handle errors (e.g., `val, _ := someFunc()`) can hide bugs and lead to unpredictable behavior.  
✅ **Fix**: Always handle errors explicitly for cleaner and more reliable code.

---

### 64.2 Misuse of Goroutines & Channels
Spawning goroutines without proper synchronization (e.g., `sync.WaitGroup`) or mismanaging channels can lead to goroutine leaks and race conditions.  
✅ **Fix**: Use `sync.WaitGroup`, context cancellation, and proper channel handling.

---

### 64.3 Variable Shadowing
Shadowing variables—especially inside loops or after using `:=`—can cause unexpected behavior by masking outer variables.  
✅ **Fix**: Avoid redeclaration inside nested scopes; prefer explicit variable naming.

---

### 64.4 Deferring Resource Cleanup in Loops
Using `defer` inside loops can delay cleanup until the function exits, consuming unnecessary resources.  
✅ **Fix**: Move `defer` into helper functions so cleanup runs per iteration.

---

### 64.5 Overusing Interfaces
Defining interfaces when only one type implements them adds unnecessary abstraction and complexity.  
✅ **Fix**: Use interfaces sparingly—mainly when multiple implementations or mocks are needed.

---

### 64.6 Incorrect Slice Handling
Slices share an underlying array; modifying one slice may affect others unintentionally.  
✅ **Fix**: Use `copy()` when separate underlying arrays are required.

---

### 64.7 Overusing `init()` Functions
While `init()` is convenient for setup, excessive use may obscure logic flow and hinder testability.  
✅ **Fix**: Prefer explicit initializers or constructors.

---

### 64.8 Range Loop Gotchas
- Using a single variable in `range` yields the **index**, not the value.  
- Variables declared via `:=` inside `range` can shadow outer ones.  
✅ **Fix**: Always check whether you need `index`, `value`, or both.

---

### 64.9 HTTP Client Without Timeout
Using Go’s default `http.Client` without setting timeouts may allow requests to hang indefinitely.  
✅ **Fix**: Always configure `Timeout` or use context-based deadlines.

---

### 64.10 Not Closing DB Rows / Misconfiguring DB Connections
Failing to `rows.Close()` prevents releasing connections back to the pool, risking exhaustion.  
✅ **Fix**: Always `defer rows.Close()` and tune connection pool settings (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`).

---

### 64.11 Goroutine Leaks
Unclosed channels or mismanaged goroutines can cause leaks.  
✅ **Fix**: Use tools like Uber’s `goleak` and `net/http/pprof` to detect leaks.

---

### 64.12 Summary Table

| **Pitfall**                          | **Why It Matters**                                | **Tip**                                            |
|-------------------------------------|---------------------------------------------------|----------------------------------------------------|
| Ignoring errors                      | Masks real bugs                                   | Always check `err`                                |
| Misuse of goroutines/channels        | Leads to races, deadlocks, leaks                  | Use `WaitGroup`, context, and proper channels      |
| Variable shadowing                   | Causes unexpected logic behavior                  | Avoid redeclaration in nested scopes               |
| Defer in loops                       | Delays cleanup until function exits               | Use helper functions for cleanup                   |
| Overusing interfaces                 | Adds unnecessary complexity                        | Use only when multiple implementations are needed  |
| Unsafe slice usage                   | Unintended shared modifications                   | Copy slices when needed                            |
| Excessive `init()` usage             | Obscures initialization code                      | Prefer explicit constructors                       |
| Range loop misuse                    | Value vs index confusion                          | Use both variables or underscore to ignore index   |
| Missing HTTP timeouts                | Potential for hanging requests                    | Set `Timeout` or use context                       |
| Not closing DB rows                  | Exhausted DB connections                          | Always `defer rows.Close()`                        |
| Goroutine leaks                      | Memory/performance degradation                    | Use leak detectors and pprof                       |


### 65. 3 ways to stop goroutines:

- `Done channel` → Send/close a channel to signal the goroutine to exit gracefully.

- `Context` (context.WithCancel/Timeout) → Standard Go way to cancel goroutines, with optional deadlines/timeouts.

- `Quit + Done channels` → Send a stop signal via one channel and wait for acknowledgment via another (bi-directional control).




---


# Others

  ## major updates in Go 1.25 – August 12, 2025

| Area                 | Highlights |
|----------------------|------------|
| **Language**         | Removed the concept of "core types" in favor of clearer prose in the spec :contentReference[oaicite:0]{index=0} |
| **Toolchain**        | `go build -asan` now enables leak detection by default, reporting memory leaks from C—even after program exit (can be disabled with `ASAN_OPTIONS=detect_leaks=0`) :contentReference[oaicite:1]{index=1} |
| **Runtime & GC**     | Experimental enhancements: new garbage collector, runtime improvements including smarter `GOMAXPROCS` behavior :contentReference[oaicite:2]{index=2} |
| **`encoding/json/v2`** | New experimental JSON API is available under the `GOEXPERIMENT=jsonv2` flag :contentReference[oaicite:3]{index=3} |
| **Debug Info**       | DWARF v5 format now used by default for debug information :contentReference[oaicite:4]{index=4} |
| **Testing**          | Added stable `testing/synctest` package for testing concurrent code :contentReference[oaicite:5]{index=5} |
| **HTTP Enhancements**| Introduced `net/http.CrossOriginProtection` for CSRF defense—no tokens or cookies required :contentReference[oaicite:6]{index=6} |
| **Sync Convenience** | Added `sync.WaitGroup.Go` as a shortcut for launching goroutines tied to `WaitGroup` :contentReference[oaicite:7]{index=7} |
| **Flight Recorder**  | Flight recorder API introduced for lightweight execution tracing :contentReference[oaicite:8]{index=8} |

