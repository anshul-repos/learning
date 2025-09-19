## Table of Contents
- [Compute Services](#compute-services)
  - [Amazon EC2](#1-amazon-ec2-elastic-cloud-compute)
  - [Amazon Lambda](#2-amazon-lambda)
  - [Amazon ECS / EKS](#3-amazon-ecs--eks)
  - [AWS Fargate](#4-aws-fargate)
  - [Amazon Elastic Beanstalk](#5-amazon-elastic-beanstalk)
- [Database Services](#database-services)
  - [Amazon RDS](#6-amazon-rds-relational-database-service)
  - [DynamoDB](#7-dynamodb)
  - [Amazon Aurora](#8-amazon-aurora)
- [Storage Services](#storage-services)
  - [Amazon S3](#9-amazon-s3-simple-storage-service)
  - [Amazon EBS](#10-amazon-ebs-elastic-block-store)
  - [Amazon EFS](#11-amazon-efs-elastic-file-system)



## Compute Services

### 1. Amazon EC2 (Elastic Cloud Compute)

Provides resizable compute capacity through virtual machines.

Features:

Multiple instance types (compute-, memory-, storage-optimized).

Choose AMIs with desired OS.

Integrates with Auto Scaling, Elastic Load Balancing, and EBS.

Use Case: Deploy a Linux EC2 instance to run a LAMP stack. 


### 2. Amazon Lambda

Fully serverless compute service; runs code in response to events/triggers.

Features:

Supports runtimes: Node.js, Python, Java, Go, etc.

Trigger via S3 uploads, API Gateway, SNS.

Pay per request and execution duration.

Use Case: Process data automatically when uploaded to S3 using a Python function. 


### 3. Amazon ECS / EKS

Container orchestration services:

ECS: Native Docker service integrated with AWS.

EKS: Managed Kubernetes environment.

Use Case: Deploy containerized microservices architecture. 


### 4. AWS Fargate

Serverless compute for containers on ECS or EKS.

Advantages:

No EC2 instance provisioning.

Automatic scaling of container workloads.

Use Case: Run stateless APIs in Docker containers using ECS + Fargate. 


### 5. Amazon Elastic Beanstalk

Simplified deployment and management of web apps.

Features:

Supports Java, Python, Node.js, .NET, PHP, Ruby, Go, Docker.

Manages provisioning, load balancing, autoscaling, and health monitoring.

Use Case: Deploy a Django application through ZIP upload. 


## Database Services

### 6. Amazon RDS (Relational Database Service)

Managed relational database supporting PostgreSQL, MariaDB, MySQL, etc.

Features:

Automated backups, replication, Multi-AZ, read replicas.

Monitoring via CloudWatch and performance insights.

Use Case: Multi-AZ PostgreSQL deployment for a SaaS backend. 


### 7. DynamoDB

Serverless NoSQL (key-value/document) database.

Features:

In-memory caching with DAX.

On-demand or provisioned throughput; auto-scaling; encryption.

Handles massive scale: up to 10 trillion requests per day, 20 million/sec. 


Use Case: Real-time gaming leaderboard. 


### 8. Amazon Aurora

High-performance, MySQL- and PostgreSQL-compatible relational DB.

Features:

Distributed, fault-tolerant, self-healing storage.

Auto-scaling, global databases, up to 15 read replicas.

Use Case: Backend DB for high-traffic e-commerce platform. 


## Storage Services

### 9. Amazon S3 (Simple Storage Service)

Highly durable object storage service.

Features:

Versioning, lifecycle policies, cross-region replication.

11 nines of durability (99.999999999%).

Use Case: Static website assets or data lake storage. 


### 10. Amazon EBS (Elastic Block Store)

Persistent block storage for EC2.

Features:

SSD (gp3, io2) and HDD (st1, sc1) volume types.

Encryption and point-in-time snapshots.

Use Case: EC2-attached storage for MySQL database. 


### 11. Amazon EFS (Elastic File System)

Scalable, serverless NFS-based file storage.

Features:

No provisioning; pay only for storage used.

NFS-compatible, accessible by multiple instances.