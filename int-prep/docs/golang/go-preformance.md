    1. What are performance metrics: 
        ► CPU usage: how my cpu used by an application
        ► Memory usage: how much memory used by application
        ► Throughput: Measures the amount of work the system completes in a given time, e.g., the number of requests handled per second.
        ► Garbage collection : time and frequency of garbage collection
        ► Latency: time taken to process a request or complete an operation
        
    
    Tools: pprof


    2. How can we improve performance of GO code?
        ► Use of int instead of string in Map if possible
        ► Sort struct fields : 
            · order of the fields used in the struct directly affects your memory usage.
            · We can auto sort struct using tools like : fieldalignment.
        ► Use of go routines
        ► Use of tools (pprof) to check the bottlenecks in system
        ► Properly use DS and Algos
        ► Use of protocol buffers instead of json for data transmission over network
        ► Pass large structs by pointer instead of pass by value
        ► Pre-allocation of slices instead of re-allocation:
            · Example
                □ Creating a slice without pre-allocation example: var numbers []int
                □ Pre-allocating a slice with capacity of 1000 example:  numbers := make([]int, 0, 1000) 
            · Every time an element is appended, the capacity is increased as needed. When the capacity is exceeded, Go automatically allocates a new array with double the previous capacity and copies the existing elements to the new array. This process can lead to unnecessary memory allocations and increased garbage collection overhead
        ► Re use object instead of reallocation