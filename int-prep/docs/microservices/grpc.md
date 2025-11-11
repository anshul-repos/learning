gRPC (Google Remote Procedure Call) is a high-performance, open-source, universal RPC (Remote Procedure Call) framework developed by Google.


Key Features:

    1. Cross-language support: gRPC is supported by most programming languages, including Go, Python, Java, and more.
    2. Efficient and fast:             Uses HTTP/2 for transport, providing multiplexing and binary framing, making it faster than REST over HTTP/1.1.
    3. Streaming support:          Supports various streaming patterns (unary, server-side, client-side, and bidirectional streaming).
    4. Code generation:             Automatically generates client and server code using .proto files.
    5. Bi-directional communication: Real-time data exchange is seamless with HTTP/2.
    6. Built-in authentication:  Supports TLS/SSL for secure communication.


gRPC Use Cases:

    1. Microservices Communication
    2. Real-time Data Streaming
    3. Polyglot Systems
    4. APIs for High Throughput Systems
    
Key Concepts in gRPC:

    1. Protocol Buffers: The serialization/deserialization mechanism. It defines:
        ○ Messages (data structures)
        ○ Services (methods)
    2. gRPC Service: A definition in .proto file specifying the available RPC methods.
    3. gRPC Stub: Auto-generated code for clients to call services and servers to implement them.
    4. gRPC Communication Types:
        a. Unary RPC: One request, one response.
        b. Server Streaming RPC: One request, multiple responses streamed from the server.
        c. Client Streaming RPC: Multiple requests streamed from the client, one response.
        d. Bidirectional Streaming RPC: Both client and server stream requests and responses simultaneously.

Common Interview Questions:

    1. What is the difference between gRPC and REST?
        ○ gRPC uses binary (Protocol Buffers), while REST uses JSON (text).
        ○ gRPC is faster due to HTTP/2 support.
        ○ gRPC supports streaming, while REST doesn’t natively.

    2. How does gRPC achieve high performance?
        ○ Efficient serialization (protobuf)
        ○ HTTP/2 (multiplexing, binary framing)

    3. How do you secure a gRPC service?
        ○ Use TLS for transport-layer security.
        ○ Implement authentication using tokens (e.g., OAuth2).

    4. What are the different types of RPC in gRPC?
        ○ Unary, Server-side streaming, Client-side streaming, and Bidirectional streaming.
    5. What are error codes in grpc?

#### gRPC Standard Error Codes

gRPC defines a standard set of error codes for all languages and platforms.  
These codes are used to indicate the outcome of an RPC call.

| Code | Name               | Description                                                                 |
|------|--------------------|-----------------------------------------------------------------------------|
| `0`  | **OK**             | Success. No error.                                                          |
| `1`  | **Canceled**       | The operation was canceled (by client or server).                           |
| `2`  | **Unknown**        | Unknown error (e.g., uncaught exception).                                   |
| `3`  | **InvalidArgument**| Client specified an invalid argument.                                       |
| `4`  | **DeadlineExceeded**| Operation took too long (timeout).                                         |
| `5`  | **NotFound**       | Requested resource not found.                                               |
| `6`  | **AlreadyExists**  | Resource already exists (e.g., duplicate entry).                            |
| `7`  | **PermissionDenied**| No permission, but *not* due to authentication failure.                    |
| `8`  | **ResourceExhausted**| Resource limits reached (e.g., quota, rate-limiting, memory).              |
| `9`  | **FailedPrecondition**| System not in required state (e.g., deleting non-empty directory).        |
| `10` | **Aborted**        | Operation aborted (e.g., concurrency conflict, transaction rollback).       |
| `11` | **OutOfRange**     | Request is out of valid range (e.g., seek past end of file).                |
| `12` | **Unimplemented**  | Method not implemented or not supported.                                    |
| `13` | **Internal**       | Internal server error (something went wrong).                               |
| `14` | **Unavailable**    | Service unavailable (e.g., server down, transient failure).                 |
| `15` | **DataLoss**       | Irrecoverable data loss or corruption.                                      |
| `16` | **Unauthenticated**| Authentication failed (invalid token, missing credentials).                 |


