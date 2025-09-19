# AWS Storage Services Overview

AWS offers a variety of storage options tailored to different workloads, from **block storage** to **object storage** and **caching solutions**.  
This document summarizes each storage type, the AWS services that support it, and real-world examples.

---

## 📊 Storage Types & Services

| **Storage Type** | **Optimized For**                                                                 | **AWS Services / Tools**                                                                 |
|------------------|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------|
| **Block**        | Applications requiring low-latency, high-performance durable storage attached to single Amazon EC2 instances or containers (e.g., databases). | - Amazon EBS <br> - Amazon EC2 Instance Store |
| **File System**  | Shared read/write access across multiple EC2 instances or containers, or from multiple on-premises servers. Ideal for team collaboration, enterprise apps, analytics, and ML training. | - Amazon EFS <br> - Amazon FSx <br> - Amazon S3 File Gateway |
| **Object**       | Read-heavy workloads like content distribution, web hosting, analytics, and ML workflows. Suited for global access over the internet. | - Amazon S3 |
| **Cache**        | High-speed temporary storage layer for frequently accessed/recently used data to reduce latency and database load. | - Amazon File Cache <br> - Amazon ElastiCache |
| **Hybrid/Edge**  | Low-latency delivery and cloud-backed storage for on-premises applications. Ensures data mobility between on-prem and cloud. | - AWS Storage Gateway (Tape Gateway, Volume Gateway) |

---

## 📦 Detailed Service Use Cases

### 1. **Amazon S3 (Simple Storage Service)**
- **Use Case**: Media storage at scale (images, videos, documents).
- **Example**:  
  - A social media platform like Instagram stores user-uploaded images and videos.  
  - An image might be stored as:  
    `https://myapp-bucket.s3.amazonaws.com/user123/photo1.jpg`

---

### 2. **Amazon EBS (Elastic Block Store)**
- **Use Case**: Persistent storage for EC2 databases and applications.
- **Example**:  
  - A company runs MySQL on EC2 and attaches a **500GB SSD (gp3) EBS volume**.  
  - Data remains intact even if the EC2 instance is stopped and later restarted.

---

### 3. **Amazon EFS (Elastic File System)**
- **Use Case**: Shared file storage across multiple EC2 instances.
- **Example**:  
  - A dev team working on a CMS mounts the same file system across EC2 instances.  
  - Scales automatically and supports thousands of concurrent connections.

---

### 4. **Amazon FSx**
- **Use Case**: Specialized file systems like Windows File Server, Lustre, NetApp ONTAP.
- **Example**:  
  - A financial company uses **Amazon FSx for Windows File Server** to store sensitive reports.  
  - Integrated with Active Directory for access permissions.

---

### 5. **Amazon S3 Glacier**
- **Use Case**: Long-term archival of compliance or historical data.
- **Example**:  
  - A healthcare company archives patient records for years at low cost.  
  - Retrieval is slower but ideal for compliance workloads.

---

### 6. **Amazon DynamoDB**
- **Use Case**: NoSQL database for high-speed, real-time workloads.
- **Example**:  
  - An online gaming platform stores **player state, scores, inventory**.  
  - Handles millions of requests/second with low latency.

---

### 7. **Amazon ElastiCache**
- **Use Case**: Caching frequently accessed data for faster performance.
- **Example**:  
  - An e-commerce site uses **ElastiCache for Redis** to cache product catalog queries.  
  - Reduces DB load and improves page speed.

---

## 🛒 Real-World Example: E-Commerce Workflow

An e-commerce company might use multiple AWS storage services together:

1. **Amazon S3** – Stores product images, videos, and backups.  
2. **Amazon EC2 + EBS** – Runs the OS and transactional databases.  
3. **Amazon RDS (MySQL)** – Manages orders and customer transactions.  
4. **Amazon DynamoDB** – Handles shopping cart data at high throughput.  
5. **Amazon ElastiCache (Redis)** – Caches product catalog queries.  
6. **Amazon S3 Glacier** – Stores long-term backups and compliance records.  

---

## ✅ Summary

- **Block Storage (EBS/Instance Store)** → High-performance database storage.  
- **File Storage (EFS/FSx)** → Shared access for multiple servers/apps.  
- **Object Storage (S3/Glacier)** → Scalable storage for media, backups, and analytics.  
- **Cache (ElastiCache/File Cache)** → Low-latency, high-throughput workloads.  
- **Hybrid/Edge (Storage Gateway)** → Bridges on-premises and cloud storage.  

---
