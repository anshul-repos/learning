    1. REST vs GRPC:
    
    REST	                                            GRPC
        
    HTTP 1.1 protocol	                                HTTP 2 protocol
    Json or xml for data exchange	                    Protocol buffer for serialization in binary form
    Uni-directional	                                    Bi-directional
    Used anywhere	                                    Used in microservice architecture
    A single client communicates with a single server.	one server to many clients, one client to many servers
        
    

# Difference between Microservice and Monolithic Architecture



| **Aspect**              | **Monolithic Architecture**                                                                                  | **Microservice Architecture**                                                                                          |
|--------------------------|--------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| **Definition**           | A single, unified codebase where all components are tightly integrated and run as a single application.       | A distributed system where each service is a separate, independent application that communicates with others via APIs. |
| **Deployment**           | Deployed as a single unit                                                                                   | Each service is deployed independently                                                                                 |
| **Scalability**          | Typically scaled by running multiple copies of the entire application.                                       | Scaled independently; only the services needing more resources are scaled.                                             |
| **Development Complexity** | Easier to develop initially                                                                                 | More complex to develop initially but easy to maintain after a period                                                   |
| **Technology Stack**     | Typically uses a single technology stack (e.g., one programming language, one database).                     | Allows for different services to use different technology stacks (polyglot programming, multiple databases).           |
| **Communication**        | Intra-process communication; function calls within the same application.                                     | Inter-process communication; typically uses HTTP/REST, gRPC, message queues, etc.                                      |
| **Fault Isolation**      | Faults can affect the entire application, making it harder to isolate and recover from failures.              | Faults in one service typically do not affect others, allowing for better fault tolerance and recovery.                 |
| **Data Management**      | Centralized database.                                                                                        | Each service manages its own database                                                                                   |
| **Development Teams**    | Typically managed by a single team due to the centralized codebase.                                          | Allows for small, autonomous teams to develop and manage individual services.                                           |
| **Testing**              | Requires end-to-end testing of the entire application.                                                       | Allows for testing individual services independently, but end-to-end testing across services can be complex.            |
| **Maintenance**          | Can become complex over time as the codebase grows and changes need to be managed carefully.                 | Easier to maintain individual services, but requires careful management of dependencies and versioning.                 |
| **Performance**          | Lower latency due to tight integration but may struggle under high load.                                     | Higher latency due to network calls between services but is better suited for high load through scaling.                |
| **Use Cases**            | Suitable for smaller, simpler applications.                                                                  | Ideal for large, complex, and evolving systems.                                                                         |



# Error Codes

    200 OK → Success

    201 Created → Resource created successfully

    400 Bad Request → Client sent invalid request

    401 Unauthorized → Authentication required

    403 Forbidden → You don’t have permission

    404 Not Found → Resource doesn’t exist

    500 Internal Server Error → Server crashed or misconfigured

    503 Service Unavailable → Server overloaded or under maintenance