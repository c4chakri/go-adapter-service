# 🚀 Go API Service – Entity Instances Proxy

A lightweight, scalable **Go-based API service** that acts as a proxy for an external Entity Instances Service. It enables dynamic schema-based instance creation with secure **Bearer token authentication** and structured relationship handling.

---

## 📖 Introduction

This service is designed to simplify interaction with an Entity Instances backend by providing a clean, secure, and extensible API layer. It handles request forwarding, authentication propagation, and structured payload management for schema-driven systems.

---

## 📚 Table of Contents

- [Features](#-features)
- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
- [Running in Production](#-running-in-production)
- [Configuration](#-configuration)
- [Project Structure](#-project-structure)
- [API Endpoints](#-api-endpoints)
- [Example Request](#-example-request)
- [Example Response](#-example-response)
- [Error Handling](#-error-handling)
- [Health Check](#-health-check)
- [Use Cases](#-use-cases)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)
- [License](#-license)

---

## ✨ Features

- 🔐 **Bearer Token Forwarding** – Securely forwards client authorization tokens
- 🧩 **Dynamic Schema Instance Creation** – Supports entities and relationships
- 🌐 **RESTful API Design** – Clean and predictable endpoints
- 🧪 **Demo Endpoints** – Built-in APIs for quick testing
- 🔄 **CORS Enabled** – Cross-origin request support
- ❤️ **Health Monitoring** – Simple health check endpoint
- ⚡ **High Performance** – Powered by Gin framework

---

## 📋 Prerequisites

- Go **1.21+**
- Git (optional)
- API testing tool (Postman / curl)

---

## 🛠️ Installation & Setup

\`\`\`bash
git clone https://github.com/c4chakri/go-adapter-service.git
cd go-api-service
go mod init go-api-service
go mod tidy
go run main.go
\`\`\`

---

## ▶️ Running in Production

\`\`\`bash
go build -o app
./app
\`\`\`

---

## ⚙️ Configuration

| Variable | Description               | Default |
|----------|--------------------------|---------|
| PORT     | Server port              | 3000    |
| GIN_MODE | Mode (debug/release)     | debug   |
| BASE_URL | External service URL     | —       |

\`\`\`bash
export PORT=3000
export GIN_MODE=release
export BASE_URL=https://api.example.com
\`\`\`

---

## 📁 Project Structure

\`\`\`
go-api-service/
├── main.go
├── handlers/
├── client/
├── models/
├── routes/
└── go.mod
\`\`\`

---

## 📡 API Endpoints

### 🔹 Create Schema Instances

POST /api/v1/schemas/:schemaId/instances

### 🔹 Get Static Entities

GET /api/v1/entities/static

### 🔹 System Stats

GET /api/v1/stats

### 🔹 Health Check

GET /api/v1/health

### 🔹 Demo Processes

GET /api/v1/processes/demo

---

## 🧪 Example Request

\`\`\`bash
curl --location 'http://localhost:3000/api/v1/schemas/{schemaId}/instances' \\
--header 'Content-Type: application/json' \\
--header 'Authorization: Bearer <your-token>' \\
--data '{
  "ontology": true,
  "data": [
    {
      "labelName": "IngestedData",
      "description": "businessEntity definition for Ingested Data"
    }
  ],
  "relationships": []
}'
\`\`\`

---

## ✅ Example Response

\`\`\`json
{
  "status": "success",
  "message": "Schema instances created successfully"
}
\`\`\`

---

## ❌ Error Handling

| Status Code | Description           |
|-------------|----------------------|
| 400         | Bad Request          |
| 401         | Unauthorized         |
| 500         | Internal Server Error|

---

## 📊 Health Check

GET /api/v1/health

\`\`\`json
{
  "status": "ok"
}
\`\`\`

---

## 🧠 Use Cases

- Proxy layer for external entity services  
- Middleware for schema-driven architectures  
- Backend for ontology-based platforms  
- Integration layer for data ingestion pipelines  

---

## 🛠️ Troubleshooting

| Issue | Possible Cause | Solution |
|------|--------------|----------|
| Service not starting | Port already in use | Change PORT |
| 401 Unauthorized | Invalid token | Verify token |
| 500 errors | Backend issue | Check BASE_URL |
| CORS issues | Missing headers | Enable CORS |

---

## 🤝 Contributing

1. Fork the repository  
2. Create a feature branch  
3. Commit changes  
4. Open a Pull Request  

---

## 📄 License

MIT License
# go-adapter-service
