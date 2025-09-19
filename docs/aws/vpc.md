A Virtual Private Cloud (VPC) is an isolated network environment within the AWS cloud, allowing you to launch AWS resources in a defined virtual network. It gives full control over your networking, including IP address ranges, subnets, routing, and security settings.


Link: 

How to setup VPC, Public, Private Subnet, NAT, Internet Gateway, Route Table?

https://www.youtube.com/watch?v=43tIX7901Gs



### 1. VPC Components

a. Subnets
    • Definition: Subnets are subdivisions of a VPC's IP address range and allow you to isolate resources within the VPC.
    • Types:
        ○ Public Subnets: Subnets where instances can communicate with the internet directly, typically for resources like web servers.
        ○ Private Subnets: Subnets that don't have direct internet access, often used for databases or backend services.
    • Best Practices:
        ○ Place public-facing resources (like EC2 instances that need to serve web traffic) in public subnets.
        ○ Place sensitive resources (like databases) in private subnets to avoid direct exposure to the internet.
b. Route Tables
    • Definition: A route table contains rules (routes) that determine how network traffic is directed within your VPC and to external networks like the internet.
    • Key Points:
        ○ Each subnet must be associated with a route table.
        ○ The route table determines how traffic is routed between subnets, to on-premise networks (via VPN), or to the internet.
    • Example:
        ○ Public Route Table: Directs traffic to the Internet Gateway for public subnets.
        ○ Private Route Table: Routes traffic internally within the VPC or to a NAT Gateway for internet-bound traffic.
c. Internet Gateway (IGW)
    • Definition: An Internet Gateway is a VPC component that allows instances in public subnets to access the internet and vice versa.
    • Use Case: Attach an IGW to your VPC, and route the public subnet’s traffic through it so that EC2 instances can communicate with the internet.
    • Key Points:
        ○ Only subnets associated with a route table that has a route to the IGW can access the internet.
        ○ It is a redundant and horizontally scalable component provided by AWS.
d. NAT Gateway
    • Definition: A NAT Gateway allows instances in private subnets to access the internet (for downloading updates, accessing external APIs, etc.) without exposing them to inbound internet traffic.
    • Use Case: If your private resources need to access the internet (e.g., for updates), they can route traffic through a NAT Gateway, which resides in a public subnet.
e. Elastic IP (EIP)
    • Definition: Elastic IPs are static IP addresses that can be assigned to your EC2 instances or other resources.
    • Use Case: Typically used with resources that require a fixed IP address, like public-facing web servers, so their IP does not change after reboot.
f. Security Groups
    • Definition: Security groups act as virtual firewalls for your instances to control inbound and outbound traffic.
    • Key Points:
        ○ By default, all inbound traffic is blocked, and all outbound traffic is allowed.
        ○ You can specify IP ranges, protocols, and ports to allow or deny specific traffic.
g. Network Access Control Lists (NACLs)
    • Definition: NACLs provide an additional layer of security by acting as a stateless firewall for controlling traffic at the subnet level.
    • Key Points:
        ○ NACLs allow or deny traffic at the subnet boundary.
        ○ They can be configured to allow or deny specific IPs, protocols, and port ranges for both inbound and outbound traffic.
h. VPC Peering
    • Definition: VPC peering allows two VPCs (in the same or different AWS accounts) to communicate with each other using private IP addresses.
    • Use Case: It’s useful for linking applications across VPCs while keeping traffic secure and avoiding exposure to the internet.

### 2. VPC Networking Flow Example
Imagine you are setting up a simple web application with a frontend in a public subnet and a database in a private subnet:
    1. VPC Creation:
        ○ You create a VPC with a CIDR block (e.g., 10.0.0.0/16).
    2. Subnets Creation:
        ○ Create a public subnet (10.0.1.0/24) for the frontend web servers.
        ○ Create a private subnet (10.0.2.0/24) for the backend database.
    3. Internet Gateway:
        ○ Attach an Internet Gateway to the VPC.
        ○ Create a route in the public subnet's route table that directs traffic to the IGW (0.0.0.0/0 -> IGW).
    4. Route Table Configuration:
        ○ The public subnet’s route table will have a route pointing to the IGW for internet-bound traffic.
        ○ The private subnet’s route table will route traffic internally or to a NAT Gateway for internet-bound traffic without exposing the instances.
    5. NAT Gateway:
        ○ Place the NAT Gateway in the public subnet and update the private subnet's route table to send outbound traffic through the NAT Gateway.
    6. Security Groups:
        ○ Assign a security group to the EC2 instance in the public subnet, allowing inbound HTTP and HTTPS traffic (e.g., 80, 443) from the internet.
        ○ Assign a security group to the database instance in the private subnet, allowing inbound traffic only from the public EC2 instance's IP address or the frontend security group.
    7. NACL Configuration (Optional):
        ○ Apply NACLs to the subnets if you want to control traffic at the subnet level in addition to the instance-level security groups.

###  Understanding AWS VPC (Virtual Private Cloud)

A Virtual Private Cloud (VPC) is an isolated network environment within the AWS cloud, allowing you to launch AWS resources in a defined virtual network. It gives full control over your networking, including IP address ranges, subnets, routing, and security settings.
Here’s a breakdown of the core components of a VPC and how they fit together to form the basic networking structure in AWS:

### 1. VPC Components

Subnets

    Definition: Subnets are subdivisions of a VPC's IP address range and allow you to isolate resources within the VPC.

    Types:
    
    Public Subnets: Subnets where instances can communicate with the internet directly, typically for resources like web servers.
    
    Private Subnets: Subnets that don't have direct internet access, often used for databases or backend services.

    Best Practices:
    Place public-facing resources (like EC2 instances that need to serve web traffic) in public subnets. 
    
    Place sensitive resources (like databases) in private subnets to avoid direct exposure to the internet.

Route Tables:

    • Definition: A route table contains rules (routes) that determine how network traffic is directed within your VPC and to external networks like the internet.

    Key Points:
    ○ Each subnet must be associated with a route table.
    ○ The route table determines how traffic is routed between subnets, to on-premise networks (via VPN), or to the internet.

    Example:
    ○ Public Route Table: Directs traffic to the Internet Gateway for public subnets.
    ○ Private Route Table: Routes traffic internally within the VPC or to a NAT Gateway for internet-bound traffic.

Internet Gateway (IGW)

    Definition: An Internet Gateway is a VPC component that allows instances in public subnets to access the internet and vice versa.

    Use Case: Attach an IGW to your VPC, and route the public subnet’s traffic through it so that EC2 instances can communicate with the internet.

    Key Points:
    ○ Only subnets associated with a route table that has a route to the IGW can access the internet.
    ○ It is a redundant and horizontally scalable component provided by AWS.

NAT Gateway

    • Definition: A NAT Gateway allows instances in private subnets to access the internet (for downloading updates, accessing external APIs, etc.) without exposing them to inbound internet traffic.

    • Use Case: If your private resources need to access the internet (e.g., for updates), they can route traffic through a NAT Gateway, which resides in a public subnet.

Elastic IP (EIP)

    • Definition: Elastic IPs are static IP addresses that can be assigned to your EC2 instances or other resources.

    • Use Case: Typically used with resources that require a fixed IP address, like public-facing web servers, so their IP does not change after reboot.

Security Groups

    • Definition: Security groups act as virtual firewalls for your instances to control inbound and outbound traffic.

    • Key Points:
        ○ By default, all inbound traffic is blocked, and all outbound traffic is allowed.
        ○ You can specify IP ranges, protocols, and ports to allow or deny specific traffic.

Network Access Control Lists (NACLs)

    • Definition: NACLs provide an additional layer of security by acting as a stateless firewall for controlling traffic at the subnet level.

    • Key Points:
        ○ NACLs allow or deny traffic at the subnet boundary.
        ○ They can be configured to allow or deny specific IPs, protocols, and port ranges for both inbound and outbound traffic.

VPC Peering

    • Definition: VPC peering allows two VPCs (in the same or different AWS accounts) to communicate with each other using private IP addresses.

    • Use Case: It’s useful for linking applications across VPCs while keeping traffic secure and avoiding exposure to the internet.

### 2. VPC Networking Flow Example
Imagine you are setting up a simple web application with a frontend in a public subnet and a database in a private subnet:

    1. VPC Creation:
        ○ You create a VPC with a CIDR block (e.g., 10.0.0.0/16).

    2. Subnets Creation:
        ○ Create a public subnet (10.0.1.0/24) for the frontend web servers.
        ○ Create a private subnet (10.0.2.0/24) for the backend database.

    3. Internet Gateway:
        ○ Attach an Internet Gateway to the VPC.
        ○ Create a route in the public subnet's route table that directs traffic to the IGW (0.0.0.0/0 -> IGW).

    4. Route Table Configuration:
        ○ The public subnet’s route table will have a route pointing to the IGW for internet-bound traffic.
        ○ The private subnet’s route table will route traffic internally or to a NAT Gateway for internet-bound traffic without exposing the instances.

    5. NAT Gateway:
        ○ Place the NAT Gateway in the public subnet and update the private subnet's route table to send outbound traffic through the NAT Gateway.

    6. Security Groups:
        ○ Assign a security group to the EC2 instance in the public subnet, allowing inbound HTTP and HTTPS traffic (e.g., 80, 443) from the internet.
        ○ Assign a security group to the database instance in the private subnet, allowing inbound traffic only from the public EC2 instance's IP address or the frontend security group.

    7. NACL Configuration (Optional):
        ○ Apply NACLs to the subnets if you want to control traffic at the subnet level in addition to the instance-level security groups.

### 3. Best Practices for AWS VPC Design

    1. Use Multiple Availability Zones (AZs):

        Spread subnets across multiple AZs for high availability and fault tolerance.

    2. Private Subnets for Sensitive Data:

        Always place sensitive resources, like databases, in private subnets to minimize exposure to the internet.

    3. Least Privilege Principle:

        Follow the least privilege principle by allowing only the necessary traffic between resources using security groups and NACLs.

    4. Use Flow Logs:
        Enable VPC flow logs to capture information about the IP traffic going to and from network interfaces in your VPC for monitoring and troubleshooting.

    5. Isolation of Environments:

        Create separate VPCs or use subnets to isolate environments like dev, staging, and production.

###  4. Tools for Managing and Automating VPC Configuration

    • AWS CloudFormation: Automate the creation of VPCs, subnets, routing tables, and more through Infrastructure as Code (IaC).

    • Terraform: Another IaC tool that allows for more flexibility across cloud providers, including AWS.

    • VPC Traffic Mirroring: Capture and inspect network traffic from EC2 instances within your VPC.
