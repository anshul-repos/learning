# 📘 System Design & Microservices Notes

A structured reference for:
- Designing a **Parking Lot (Vehicle Parking Space Management) System**
- **Adding a new endpoint** in a microservice architecture
- **Improving API performance**

---

## 1) 🅿️ Parking Lot System — Iteration 1 Requirements

Design and implement software for **Vehicle Parking Space Management** with the following requirements:

### ✅ Dynamic Sizing
- The software must **scale from small to large** parking lots.

### 🗺️ Parking Lot Map Input
- Supports loading a **parking lot map** (e.g., `HashMap` / `Dictionary`) as input.
- The map contains:
  - A **unique number** (ID) for each parking space designed by the engineer.
  - **Type of vehicle** supported for that space (e.g., two-wheeler, hatchback, SUV, large).

### 📊 Real-Time Summary
Provide functionality to get real-time stats:
- **Total vs Occupied** spaces
- **Two Wheeler** spaces
- **Hatchback (sub 4m)** spaces
- **SUV (4–5m)** spaces
- **Large vehicles (>5m)** spaces

### 🎯 Allotment Logic
- Allocate a **best-fit** parking space for an entering vehicle in real time.
- Once allotted, **mark space as occupied**; prevent re-allocation until released.
- If no suitable space is available, **return a proper message**.

### 🔁 Release Flow
- When a vehicle departs, **release the associated space** and make it available again.

> 💡 *Implementation guidance (OOP, Patterns, Clean Code):*  
> Identify key objects (e.g., `ParkingLot`, `ParkingSpace`, `Vehicle`, `AllocationStrategy`, `SummaryService`), define relationships and behaviors.  
> Consider patterns like **Strategy** (best-fit allocation), **Factory** (vehicle/space creation), **Repository** (map-backed persistence), and **Observer** (real-time summary updates).

---

## 2) 🔌 Steps to Add a New Endpoint in Microservice Architecture

#### 1. Understand the Requirement
- Clarify the functionality.
- Define **request/response schema**.
- Decide whether it fits an **existing service** or needs a **new microservice**.

#### 2. Update API Contracts / Specification
- **REST (OpenAPI/Swagger)** or **gRPC (protobuf)**: add the endpoint definition.

**Example — REST (Swagger/OpenAPI YAML):**
```yaml
/users/{id}/profile:
  get:
    summary: Get user profile
    parameters:
      - in: path
        name: id
        required: true
        schema:
          type: string
    responses:
      '200':
        description: Profile data
```

**Example — gRPC (.proto):**

    rpc GetUserProfile(GetUserRequest) returns (UserProfileResponse);

#### 3. Implement the Endpoint in the Service

- **Add a handler/controller.**

Implement business logic (DB/cache/calls to other services).

- **Example — Go HTTP handler:**

```yaml
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    profile, err := h.service.GetProfile(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(profile)
}
```

### Adding a New Endpoint in Microservice Architecture

#### 4. Service-to-Service Communication (if needed)
- Integrate via **REST / gRPC / Kafka**.  
- Use **client SDKs** and **service discovery** (Kubernetes Service, Consul, Istio).  

---

#### 5. Update API Gateway / Ingress Rules
- Register the new route at **API Gateway/Ingress**.  
- Apply **authentication, rate limiting, routing policies**.  

---

#### 6. Update Configuration / Service Registry
- Ensure the service is discoverable with the new route.  
  - Consul  
  - Eureka  
  - Kubernetes  

---

#### 7. Testing
- **Unit tests** for handler/service logic.  
- **Integration tests** (mock downstreams if needed).  
- **E2E tests** in staging.  

---

#### 8. Observability
- Add **logging, metrics, tracing**.  
  - Prometheus  
  - OpenTelemetry  
  - ELK Stack  
- Update **dashboards/alerts**.  

---

#### 9. CI/CD Deployment
- **Commit and push code**.  
- **CI**: linting, unit, integration tests.  
- **CD**: deploy to staging → production.  

---

#### 10. Documentation & Communication
- Update API docs:  
  - Swagger UI  
  - Postman Collections  
  - gRPC Docs  
- Notify **consumers/teams** about the new endpoint.  

---

## My Approach (Summary)
1. Update the API contract (**Swagger/proto**).  
2. Implement **controller → service → repository layers**.  
3. Integrate downstream services via **gRPC/REST** with service discovery.  
4. Expose via **API Gateway/Ingress**.  
5. Add **tests & observability**.  
6. Ship through **CI/CD pipelines**.  
7. Update **documentation** and notify stakeholders.  


## 3) How to improve API Performance ? 

    1. Pagination
    2. Async Logging
    3. Data Caching
    4. Payload Compression
    5. Connection Pool
    

    • Result Pagination: 
        This method is used to optimize large result sets by streaming them back to the client, enhancing service responsiveness and user experience. 

    • Asynchronous Logging: 
        This approach involves sending logs to a lock-free buffer and returning immediately, rather than dealing with the disk on every call. Logs are periodically flushed to the disk, significantly reducing I/O overhead. 

    • Data Caching: 
        Frequently accessed data can be stored in a cache to speed up retrieval. Clients check the cache before querying the database, with data storage solutions like Redis offering faster access due to in-memory storage. 

    • Payload Compression: 
        To reduce data transmission time, requests and responses can be compressed (e.g., using gzip), making the upload and download processes quicker. 

    • Connection Pooling: 
        This technique involves using a pool of open connections to manage database interaction, which reduces the overhead associated with opening and closing connections each time data needs to be loaded. The pool manages the lifecycle of connections for efficient resource use.



## 3)    SQL vs NoSQL

| No. | Criteria            | SQL                                                                 | NoSQL                                                                 |
|-----|---------------------|---------------------------------------------------------------------|----------------------------------------------------------------------|
| 1   | **Data Structure**  | Structured, predefined schema                                       | Unstructured, flexible schema                                        |
| 2   | **Scalability**     | Vertical Scaling                                                   | Horizontal Scaling                                                   |
| 3   | **Transaction Support** | Follows **ACID** properties                                         | Follows **CAP** (Consistency, Availability, Partition Tolerance)     |
| 4   | **Performance**     | Slower for large-scale data                                        | Faster for high read/write throughput, low-latency operations        |
| 5   | **Query Complexity**| Supports complex queries                                           | Not suitable for complex queries                                     |
| 6   | **Cost**            | Higher (licensing and maintenance cost)                            | Lower (open-source options available)                                |
| 7   | **Flexibility**     | Schema is rigid                                                    | Schema is flexible                                                   |
| 8   | **Use Case**        | Financial systems, ERP, CRM                                        | Big Data, real-time analytics, social networks, CMS                  |

---

✅ **Summary:**  
- **SQL** is best for structured, relational data where ACID transactions and complex queries are needed.  
- **NoSQL** is preferred for unstructured or semi-structured data, horizontal scalability, and real-time applications.  
