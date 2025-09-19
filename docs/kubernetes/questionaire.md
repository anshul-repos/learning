# 🚀 Kubernetes Interview Q&A

A collection of commonly asked Kubernetes interview questions with detailed answers.

---

## 1️⃣ Kubernetes Architecture & Components

- Control Plane (Master Node):

        API Server → Entry point (kubectl, clients)

        etcd → Key-value store for cluster state

        Controller Manager → Ensures desired state

        Scheduler → Assigns pods to nodes

- Worker Nodes:

        kubelet → Communicates with control plane, runs pods

        kube-proxy → Networking (Services, load balancing)

        Container Runtime → Docker, containerd, CRI-O

## 2️⃣ How to Get Logs from a Pod?

```bash
kubectl logs <pod-name> -n <namespace>
kubectl logs <pod-name> -c <container-name> -n <namespace>
kubectl logs <pod-name> --all-containers -n <namespace>
```

## 3️⃣ How to Monitor a Kubernetes Cluster?

- **kubectl commands:**
  ```bash
  kubectl top nodes
  kubectl top pods
  ```

- Metrics Server → Collects resource metrics (CPU, memory), required for HPA.

- Logging & Monitoring tools:

        Prometheus + Grafana → Metrics, dashboards

        ELK/EFK stack → Centralized logging

        Kube-state-metrics → Cluster state info

        Jaeger / Zipkin → Tracing

## 4️⃣ Role of Load Balancers in Kubernetes

- **Service type = LoadBalancer** → Exposes app to external world  
- **Balances traffic across Pods**  
- **Can be:**
  - Cloud provider LB (AWS ELB, GCP LB, Azure LB)  
  - External solutions (NGINX, HAProxy, MetalLB)  

---

## 5️⃣ Init Container

- Runs **before main containers**  
- **Use cases:**
  - Setting up configuration  
  - Checking dependencies (DB ready, API reachable)  
  - Pre-populating files  

**YAML Example:**
```yaml
spec:
  initContainers:
    - name: init-myservice
      image: busybox
      command: ['sh', '-c', 'echo Initializing...']
```

## 6️⃣ Ingress & Egress

**Ingress:**
- Manages **incoming traffic** (external → cluster)  
- Example: **NGINX ingress controller**  
- Routes traffic based on host/path  

**Egress:**
- Manages **outgoing traffic** (cluster → external services)  
- Controlled via **NetworkPolicies**  

---

## 7️⃣ Layering in Docker

- Each command (`FROM`, `RUN`, `COPY`) creates a **new image layer**  
- Layers are **cached & reusable** → faster builds  
- Layers are **read-only**, with a writable layer added at runtime  

---

## 8️⃣ Deployment vs StatefulSet

| Feature       | Deployment (Stateless) | StatefulSet (Stateful) |
|---------------|-------------------------|-------------------------|
| **Use Case**  | Web servers, APIs      | Databases, Kafka, Redis |
| **Pods**      | Interchangeable        | Stable identity (name, storage, network) |
| **Scaling**   | Easy scaling up/down   | Each Pod maintains its identity |
| **Storage**   | Ephemeral              | PersistentVolumeClaims bound to Pods |

---

## 9️⃣ Debugging Pod Issues

### 🔍 Useful Commands
```bash
# Check Pod Status
kubectl get pods -n <namespace>

# Describe Pod
kubectl describe pod <pod-name> -n <namespace>

# Check Logs
kubectl logs <pod-name> -c <container-name> -n <namespace>

# Exec into Pod
kubectl exec -it <pod-name> -n <namespace> -- /bin/sh

# Check Events
kubectl get events -n <namespace> --sort-by='.lastTimestamp'
```
## 🔟 Dockerfile Format & ENTRYPOINT vs CMD

**Basic Dockerfile Example:**
```dockerfile
FROM base-image
RUN apt-get install -y ...
COPY . /app
CMD ["python", "app.py"]    
```
- **CMD** → Provides default arguments for the container  
- **ENTRYPOINT** → Defines the main executable  

### ✅ Best Practices
- Use **ENTRYPOINT** for the main process  
- Use **CMD** for default arguments  

## 1️⃣1️⃣ How to Scale Pods

Manual Scaling:

```bash
kubectl scale deployment <deployment-name> --replicas=5 -n <namespace>
```
Auto Scaling (HPA):
```bash
kubectl autoscale deployment <deployment-name> --cpu-percent=80 --min=2 --max=10
```