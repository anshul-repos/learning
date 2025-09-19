# 📌 Kubernetes CRDs (Custom Resource Definitions)

**CRDs (Custom Resource Definitions)** are a feature in Kubernetes that allows users to define their own resource types.  
These custom resources can be used alongside the built-in Kubernetes resources (like **Pods, Services, Deployments**) to extend the Kubernetes API.  

CRDs provide a way to implement and manage **custom functionality** in a Kubernetes cluster.

---

## 🔑 Key Concepts of CRDs

1. **Custom Resources**  
   - An extension of the Kubernetes API that provides a new object type, similar to built-in resources.  
   - Allows storing and retrieving structured data.  
   - Examples: `BackupSchedule`, `DatabaseCluster`, `MyAppConfig`.

2. **Custom Resource Definition (CRD)**  
   - Defines a new resource type in Kubernetes.  
   - Once created, users can create and interact with instances of that custom resource via `kubectl` or the Kubernetes API.

3. **Controllers**  
   - Control loops that manage and automate the lifecycle and behavior of resources in Kubernetes.

4. **Relation Between CRD and Controllers**  
   - CRDs define **custom resource structures**.  
   - Controllers implement the **logic** and manage instances of these resources.

---

## ❓ Why Use CRDs?

CRDs are used to **extend Kubernetes** to manage applications, infrastructure, or services that are not natively supported by the Kubernetes API.  

With CRDs, you can build an **Operator** (a controller that manages these custom resources) that automates their lifecycle.

---

## 📝 Example of a CRD

Let’s say you want to define a **custom resource for `DatabaseCluster`** to manage database clusters in Kubernetes.

### Define the CRD (YAML Example)

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: databaseclusters.mycompany.com
spec:
  group: mycompany.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                size:
                  type: integer
                version:
                  type: string
  scope: Namespaced
  names:
    plural: databaseclusters
    singular: databasecluster
    kind: DatabaseCluster
    shortNames:
      - dbcluster
```


This YAML defines a custom resource DatabaseCluster in the group mycompany.com.
It includes attributes like size and version.

Creating an Instance of the Custom Resource

Once the CRD is created, you can create a custom resource of that type:

```yaml
apiVersion: mycompany.com/v1
kind: DatabaseCluster
metadata:
  name: example-cluster
spec:
  size: 3
  version: "1.0"
```

This creates a DatabaseCluster object named example-cluster with a size of 3 and version 1.0.

**Managing Custom Resources**

After defining a CRD, you can interact with custom resources just like built-in ones:

```bash
kubectl get databaseclusters
kubectl describe databasecluster example-cluster
```