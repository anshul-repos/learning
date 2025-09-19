# ✅ AZ-204: Implement Containerized Solutions – Lab Scenarios

## 🔹 Lab 1: Build and Push a Docker Image to Azure Container Registry (ACR)

**Objective:**
- Build a Docker image of a simple app
- Push it to Azure Container Registry

**Steps:**
1. Create a Go or Node.js app (e.g., `go-rest-api`, `express-app`)
2. Write a `Dockerfile`
3. Build the image:
```bash
docker build -t hello-api .
```

## 🔹 Lab 5: Use Deployment Slots in Azure App Service for Containers

**Objective:**
- Deploy the container image to Azure App Service
- Create and swap deployment slots

**Steps:**

1. **Create Web App for Containers**
```bash
az webapp create \
  --resource-group <your-rg> \
  --plan <app-service-plan> \
  --name <webapp-name> \
  --deployment-container-image-name <acr-name>.azurecr.io/hello-api:v1
```

2. **Enable ACR Authentication**
```bash
az webapp config container set \
  --name <webapp-name> \
  --resource-group <your-rg> \
  --docker-custom-image-name <acr-name>.azurecr.io/hello-api:v1 \
  --docker-registry-server-url https://<acr-name>.azurecr.io \
  --docker-registry-server-user <user> \
  --docker-registry-server-password <password>
```

3. **Add Deployment Slot**

```bash
az webapp deployment slot create \
  --name <webapp-name> \
  --resource-group <your-rg> \
  --slot staging
```


### ✅ Skills Covered: App Service, container slots, blue-green deployment