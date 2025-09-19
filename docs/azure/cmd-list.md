# ⚡ Azure CLI Cheat Sheet for AZ-204

This cheat sheet covers the **most commonly used Azure CLI commands** for **AZ-204 exam preparation** and **interview readiness**.

---

## 🔑 Authentication
```bash
# Login to Azure
az login

# Show currently signed-in account
az account show

# List available subscriptions
az account list -o table

# Set active subscription
az account set --subscription <SUBSCRIPTION_ID>
```

# List all resource groups

```bash
az group list -o table

# Create a resource group
az group create --name myResourceGroup --location eastus

# Delete a resource group
az group delete --name myResourceGroup --yes --no-wait
```

# 📦 Azure Container Registry (ACR)
```bash
# List container registries
az acr list -o table

# Create a new container registry
az acr create --resource-group myResourceGroup --name myContainerRegistry --sku Basic

# Login to ACR
az acr login --name myContainerRegistry

# List repositories inside ACR
az acr repository list --name myContainerRegistry -o table

# Delete an image from ACR
az acr repository delete --name myContainerRegistry --repository myapp --image latest
```

# 🌐 App Service (Web Apps)
```bash
# List all web apps
az webapp list -o table

# Create a web app
az webapp create --resource-group myResourceGroup --plan myAppServicePlan --name myWebApp --runtime "DOTNET:6"

# Deploy code to web app
az webapp up --name myWebApp --resource-group myResourceGroup --runtime "DOTNET:6"

# View app logs
az webapp log tail --name myWebApp --resource-group myResourceGroup
```

# ⚡ Azure Functions
```bash
# Create a Function App
az functionapp create --resource-group myResourceGroup --consumption-plan-location eastus \
  --runtime dotnet --functions-version 4 --name myFunctionApp --storage-account mystorageaccount

# List all Function Apps
az functionapp list -o table

# Stream logs from a Function App
az functionapp log streaming --name myFunctionApp --resource-group myResourceGroup
```

# 💾 Storage (Blob, Queue, Table)
```bash
# List storage accounts
az storage account list -o table

# Create a storage account
az storage account create --name mystorageaccount --resource-group myResourceGroup --location eastus --sku Standard_LRS

# Create a blob container
az storage container create --name mycontainer --account-name mystorageaccount

# Upload a file to blob
az storage blob upload --account-name mystorageaccount --container-name mycontainer \
  --name myfile.txt --file ./myfile.txt

# List blobs in container
az storage blob list --account-name mystorageaccount --container-name mycontainer -o table
```

# 🗄️ Cosmos DB
```bash
# List Cosmos DB accounts
az cosmosdb list -o table

# Create a Cosmos DB account (SQL API)
az cosmosdb create --name mycosmos --resource-group myResourceGroup --kind GlobalDocumentDB

# Create a database
az cosmosdb sql database create --account-name mycosmos --name mydb --resource-group myResourceGroup

# Create a container
az cosmosdb sql container create --account-name mycosmos --database-name mydb \
  --name mycontainer --partition-key-path "/id" --resource-group myResourceGroup
```

# 🔐 Security & Identity
```bash
# List Key Vaults
az keyvault list -o table

# Create a Key Vault
az keyvault create --name myKeyVault --resource-group myResourceGroup --location eastus

# Add a secret to Key Vault
az keyvault secret set --vault-name myKeyVault --name mySecret --value "SuperSecretValue"

# Retrieve a secret
az keyvault secret show --vault-name myKeyVault --name mySecret
```

# 📊 Monitoring & Insights
```bash
# Enable Application Insights for a web app
az monitor app-insights component create --app myAppInsights --location eastus --resource-group myResourceGroup

# List alerts
az monitor metrics alert list -g myResourceGroup

# Stream logs
az monitor activity-log list --max-events 5
```

# 📩 Messaging & Event Services
```bash
# Event Grid - list topics
az eventgrid topic list -o table

# Service Bus - list namespaces
az servicebus namespace list -o table

# Queue Storage - create a queue
az storage queue create --name myqueue --account-name mystorageaccount
```


# ✅ Pro Tips

Always add **-o** table for readable outputs.

Use **--query** to filter specific fields. 

Example:
```bash
az group list --query "[].name" -o table
```

Use **az find** "command" to get examples for any command. 

Example:
```bash
az find "az webapp"
```