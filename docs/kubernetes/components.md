# Kubernetes Objects Overview

Kubernetes (K8s) provides a rich set of objects to deploy, scale, and manage containerized applications.  
This document covers the most common Kubernetes objects and how they interact with each other.

---

## 📦 Core Workload Objects

| **Object**        | **Description**                                                                                  | **Interaction**                                                                 |
|--------------------|--------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| **Pod**           | Smallest deployable unit in Kubernetes; represents a running process in your cluster. Can have one or more containers sharing the same network namespace. | Managed by higher-level controllers like Deployments, StatefulSets, or DaemonSets. |
| **ReplicaSet**    | Ensures a specified number of Pod replicas are always running.                                    | Typically managed by a Deployment, which creates/manages ReplicaSets.            |
| **Deployment**    | Manages Pods and ReplicaSets. Supports declarative updates and rollouts.                          | Creates and manages ReplicaSets → ReplicaSets manage Pods.                       |
| **StatefulSet**   | Manages Pods requiring stable identities and persistent storage (stateful apps).                  | Similar to Deployment, but provides identity and ordering guarantees.             |
| **DaemonSet**     | Ensures one copy of a Pod runs on all (or some) nodes (e.g., monitoring/logging agents).          | Manages Pods across nodes.                                                       |
| **Job**           | Creates Pods that run to completion (batch workloads).                                            | Runs Pods until tasks complete.                                                  |
| **CronJob**       | Schedules Jobs at regular intervals (like Unix cron).                                             | Creates Jobs on a schedule → Jobs create Pods.                                   |

---

## 🌐 Networking & Access

| **Object**     | **Description**                                                                 | **Interaction**                                                                 |
|----------------|---------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| **Service**    | Provides a stable IP/DNS for Pods, load balances traffic among them.             | Routes traffic to Pods based on label selectors.                                |
| **Ingress**    | Manages external HTTP(S) access to Services. Provides routing rules.            | Works with Services to expose applications externally.                          |
| **Namespace**  | Logical partitioning of cluster resources for organization and isolation.        | Pods, Services, ConfigMaps, etc., live inside namespaces.                       |
| **NetworkPolicy** | Defines rules for traffic between Pods (allow/deny).                           | Applied to Pods to secure communication.                                        |

---

## ⚙️ Configuration & Storage

| **Object**            | **Description**                                                                          | **Interaction**                                                                 |
|------------------------|------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| **ConfigMap**         | Stores configuration data (non-sensitive) in key-value pairs.                             | Pods can consume via env variables, args, or mounted files.                     |
| **Secret**            | Stores sensitive information (passwords, keys, tokens).                                   | Pods can consume similarly to ConfigMaps but kept encrypted.                    |
| **PersistentVolume (PV)** | Represents cluster-wide storage provisioned by admin or dynamically.                     | Provides persistent storage to Pods.                                            |
| **PersistentVolumeClaim (PVC)** | User request for storage that binds to a PV.                                         | Pods use PVCs to access persistent storage.                                     |

---

## 🖼️ Interaction Diagram

```mermaid
flowchart TD
    subgraph User["User / DevOps"]
        Deploy[Deployment]
        State[StatefulSet]
        Daemon[DaemonSet]
        Job[Job]
        Cron[CronJob]
    end

    Deploy --> RS[ReplicaSet]
    RS --> Pod1[Pod]
    State --> Pod2[Stateful Pod]
    Daemon --> Pod3[Daemon Pod]
    Job --> Pod4[Job Pod]
    Cron --> Job

    subgraph Config["Config & Storage"]
        CM[ConfigMap]
        Secret[Secret]
        PVC[PersistentVolumeClaim]
        PV[PersistentVolume]
    end

    Pod1 --> CM
    Pod1 --> Secret
    Pod2 --> PVC --> PV

    subgraph Network["Networking"]
        Service[Service]
        Ingress[Ingress]
        NetPol[NetworkPolicy]
    end

    Pod1 --> Service
    Pod2 --> Service
    Pod3 --> Service
    Ingress --> Service
    NetPol --> Pod1
    NetPol --> Pod2
