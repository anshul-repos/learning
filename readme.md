This repository will be used for learning content.


# 📂 Repository Structure

- **learning/**
  - **int-prep/**
    - **coding/**
      - **general-topics/**
        - `concurrency/` → Goroutines, channels, sync primitives  
        - `errorhandling/` → Error patterns and best practices  
        - `go-loadbalancer/` → Simple load balancer implementations  
        - `oops/` → Object-oriented principles in Go  
        - `restapi/` → Building REST APIs in Go  
        - `workerpool/` → Worker pool patterns  
      - **go-algos/** → Go implementations of common algorithms  
      - **go-data-structures/**
        - `array/` → Array-based problems and patterns  
        - `string/` → String manipulation problems  
  - **docs/**
    - **aws/**
    - **azure/**
    - **cicd/**
    - **dbms/**
    - **golang/**
    - **grpc/**
    - **kafka/**
    - **kubernetes/**
    - **monitoring/**
    - **start-here/**
    - **system-design/**




# 1.  REST API CODE:
    - How to run: go run api.go (server will start running on port 8181)
    - How to send post curl request: 
            ansh@Mac golang % curl -X POST http://localhost:8181/employee -H "Content-Type: application/json" -d '{"id":1,"name":"Alice"}'{"id":1,"name":"Alice"}

            ansh@Mac golang % curl -X POST http://localhost:8181/employee -H "Content-Type: application/json" -d '{"id":2,"name":"BOB"}' 
            {"id":2,"name":"BOB"}

    - How to send GET request: http://localhost:8181/employee from browser.

