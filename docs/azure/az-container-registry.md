# 🐳 Azure Container Registry (ACR) & Containerized Solutions

## 📦 ACR Service Tiers

| Tier     | Description |
|----------|-------------|
| **Basic** | Cost-optimized entry point for developers learning ACR. Includes the same programmatic capabilities as Standard & Premium (Microsoft Entra ID auth, image deletion, webhooks). Best for **low usage scenarios** with limited storage & throughput. |
| **Standard** | Same features as Basic, but with **more storage & throughput**. Recommended for **most production workloads**. |
| **Premium** | Designed for **high-volume, enterprise workloads**. Provides maximum storage & throughput plus advanced features: <br> 🔹 **Geo-replication** – single registry across multiple regions <br> 🔹 **Content trust** – image tag signing <br> 🔹 **Private link** – restrict access via private endpoints |

---

## ⚡ ACR Tasks (Automated Image Builds)

**ACR Tasks** streamline container image workflows: build → test → push → deploy.  

## 1. Quick Task
- Build and push a **single container image** to a container registry **on-demand**.
- Does not require a local Docker Engine installation.
- Think of it as `docker build` + `docker push` in the cloud.

## 2. Automatically Triggered Tasks
Enable one or more triggers to build an image automatically:

- **Trigger on source code update**
- **Trigger on base image update**
- **Trigger on a schedule**

---

### Notes
- Azure Container Registry (ACR) is similar to a **Git repository**.
- A repository inside ACR corresponds to a **Docker image**.

---

### Scenario Comparison

| Scenario                  | Use ACI? | Use ACA? |
|----------------------------|:--------:|:--------:|
| Run a script once a day    | ✅ Yes   | ❌ Overkill |
| Deploy a scalable REST API | ❌ No    | ✅ Yes |
