### 1. Scaling:

How does cloud scaling work, and how would you design for scalability and high availability in AWS?

Types of Scaling:

    Vertical Scaling (Scaling Up/Down): Increasing or decreasing the resources (CPU, RAM) of a single instance. This can be done manually or automatically.

    Horizontal Scaling (Scaling Out/In): Adding or removing instances to handle increased or decreased load. This is more common in cloud environments due to its flexibility and redundancy.


### 2. Core Concepts:

    • EC2 (Elastic Compute Cloud): Understand how to launch and manage virtual machines.

    • S3 (Simple Storage Service): Review how S3 works for storing and retrieving objects, along with lifecycle policies.

    • VPC (Virtual Private Cloud): Understand the basic networking in AWS, including subnets, routing tables, and Internet Gateways.

    • IAM (Identity and Access Management): Study how AWS handles roles, permissions, and security policies.

    • EKS (Elastic Kubernetes Service): Review how to deploy and manage Kubernetes clusters in AWS. Understand the integration between EKS and other AWS services like VPC, IAM, and CloudWatch.

    • Elastic Load Balancing and Auto Scaling: Know how to configure AWS Load Balancers and set up auto-scaling policies for high availability.

Key Services:

        • RDS (Relational Database Service): How RDS is used for managed databases.
        • CloudWatch: AWS service for monitoring and logging.
        • Lambda: Event-driven, serverless computing platform that runs code without provisioning servers.

Hands-on Practice:

    • Deploy a sample application on an EC2 instance.
    • Set up an S3 bucket and configure permissions for public and private objects.

Launch an EKS cluster and deploy a Kubernetes application.