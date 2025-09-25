## Garbage Collector (GC)

- Golang has a built-in garbage collector that runs in the background. It periodically:
    - Scans the program to identify unused memory.
    - Frees up that unused memory so it can be reused. 
    
- The Go GC uses a generational `mark-and-sweep` algorithm, with a concurrent and tri-color marking approach to minimize pause durations and optimize performance.


`Generational Mark-and-Sweep Algorithm`

The garbage collector in Go doesn't strictly use a generational model (as some other languages do, like Java), but its behavior is optimized for short-lived and long-lived objects:

    • Short-Lived Objects: Many objects in programs are created and used quickly (e.g., temporary variables). These are collected faster because the GC prioritizes them.
    • Long-Lived Objects: Objects that stay around longer are checked less frequently, as they are less likely to become garbage.

`Tri-Color Marking Approach`

- The tri-color marking approach is a way to efficiently identify which parts of memory can be safely cleaned up. 

- It divides objects into three categories during garbage collection:

    • White: Objects that haven’t been checked yet and are candidates for deletion.
    • Gray: Objects that are being checked. These objects are still reachable but may point to other objects that need checking.
    • Black: Objects that are fully checked and still in use (i.e., reachable). These will not be collected.
    
Process:

    1. Start with all objects in white.

    2. Move reachable objects to gray and then to black as they are examined.

    3. Any objects left in white after the marking phase are unreachable and can be safely collected.


- Go's garbage collector works concurrently with your application, meaning it doesn’t completely stop the program to clean up memory. Instead, it does most of its work in parallel with your running code.




## Context

- A context in Go is a way to pass:

    1. request-scoped data

    2. cancellation signals
    
    3. deadlines across API boundaries and between goroutines.
    
- Context is commonly used in functions that perform long-running or concurrent operations, such as HTTP handlers, database queries, and more.

- Types:

    1. `context.Background()` : used in main entry point of a program where no parent context exist

    2. `context.TODO()`:             used when you’re unsure about which context to use

    3. `context.WithCancel()`:   context that can be manually canceled

    4. `context.WithTimeout()`: context will be automatically canceled when the specified duration elapses

    5. `context.WithDeadline()`: context that will be canceled at the specified time

    6. `context.WithValue()`: context that carries a key-value pair, useful for passing request-scoped data
    
---

# Middleware in Golang

## 1. Definition
Middleware in Go is a function that sits **between the incoming HTTP request and the final handler**. It allows us to perform tasks **before or after the main business logic** without cluttering the handler itself.

---

## 2. Why Middleware is Useful
- **Cross-cutting concerns:** Logging, authentication, error recovery, CORS, rate-limiting, metrics, etc.  
- **Code reuse:** Write once, use across multiple routes.  
- **Cleaner handlers:** Handlers focus only on business logic.  
- **Composable:** Multiple middlewares can be chained in a pipeline.

> Analogy: Think of it as a processing pipeline—each middleware is a station that modifies or validates the request before it reaches the final handler.

---

## 3. Go-Specific Details
- Middleware wraps `http.Handler` or `http.HandlerFunc`.  
- Can **modify requests**, **stop execution**, or **inject values into context** for downstream handlers.  

---

## 4. Examples

### Logging Middleware

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request:", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

### Auth Middleware
```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") != "secret" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```
### Recovery Middleware

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

### 5. Middleware and Dependency Injection

> Middleware can also act like a form of dependency injection because it can inject cross-cutting functionality or shared data (like DB connections, logger, or auth info) into the request context. The handler can then access this data without managing it directly.