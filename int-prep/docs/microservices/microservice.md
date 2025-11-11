# Service Registry in microservices

## Option A: Static Registration (Dev / Simple Setup)
## Option B: Dynamic Registration (Production / Kubernetes)



### Static Registration (Dev / Simple Setup)

- Gateway keeps a map of service names → gRPC addresses
- Example:

    ```go
        var services = map[string]string{
            "patient": "localhost:50051",
            "doctor":  "localhost:50052",
        }
    ```

### Dynamic Registration (Production / Kubernetes)

- Each service registers with a discovery system
- Kubernetes Service (DNS-based): `patient-service.default.svc.cluster.local:50051`
- Consul / etcd: services register themselves with name & address
- Gateway queries discovery system to find service endpoints dynamically.

