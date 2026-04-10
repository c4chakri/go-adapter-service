package handlers

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "strings"
    "go-api-service/models"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    baseURL string
}

func NewHandler() *Handler {
    return &Handler{
        baseURL: "https://ig.gov-cloud.ai/pi-entity-instances-service/v2.0/schemas/69d855b76163360be3980512/instances",
    }
}

func (h *Handler) CreateSchemaInstances(c *gin.Context) {
    var req models.SchemaRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid request body",
            "error":   err.Error(),
        })
        return
    }

    // Get the Authorization header from the incoming request
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Authorization header is required",
        })
        return
    }

    // Extract the token (remove "Bearer " prefix if present)
    token := strings.TrimPrefix(authHeader, "Bearer ")
    if token == authHeader {
        // If no "Bearer " prefix, use the header as is
        token = authHeader
    }

    // Marshal the request body to JSON
    jsonData, err := json.Marshal(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to marshal request data",
            "error":   err.Error(),
        })
        return
    }

    // Create HTTP request to external API
    httpReq, err := http.NewRequest("POST", h.baseURL, bytes.NewBuffer(jsonData))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to create HTTP request",
            "error":   err.Error(),
        })
        return
    }

    // Set headers using the token from the incoming request
    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Authorization", "Bearer "+token)

    // Make the request
    client := &http.Client{}
    resp, err := client.Do(httpReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to make request to external API",
            "error":   err.Error(),
        })
        return
    }
    defer resp.Body.Close()

    // Read response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Failed to read response from external API",
            "error":   err.Error(),
        })
        return
    }

    // Set the content type header from the external API response
    contentType := resp.Header.Get("Content-Type")
    if contentType != "" {
        c.Header("Content-Type", contentType)
    }

    // Return the response with the same status code
    c.Data(resp.StatusCode, contentType, body)
}

func (h *Handler) GetStaticEntities(c *gin.Context) {
    staticData := models.StaticDataResponse{
        Entities: []models.EntityInstance{
            {LabelName: "IngestedData", Description: "businessEntity definition for Ingested Data"},
            {LabelName: "BusinessRules", Description: "businessEntity definition for Business Rules"},
            {LabelName: "PricingParameters", Description: "businessEntity definition for Pricing Parameters"},
            {LabelName: "ModelConfiguration", Description: "businessEntity definition for Model Configuration"},
            {LabelName: "MonetizationModel", Description: "businessEntity definition for Monetization Model"},
        },
        Stats: models.SystemStats{
            TotalEntities:     15,
            TotalRelations:    14,
            ProcessCount:      1,
            BusinessRuleCount: 2,
        },
    }

    c.JSON(http.StatusOK, models.DemoResponse{
        Status:  "success",
        Message: "Static entities retrieved successfully",
        Data:    staticData,
    })
}

func (h *Handler) GetSystemStats(c *gin.Context) {
    stats := models.SystemStats{
        TotalEntities:     15,
        TotalRelations:    14,
        ProcessCount:      1,
        BusinessRuleCount: 2,
    }

    c.JSON(http.StatusOK, models.DemoResponse{
        Status:  "success",
        Message: "System statistics retrieved successfully",
        Data:    stats,
    })
}

func (h *Handler) GetHealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":    "healthy",
        "service":   "go-api-service",
        "version":   "1.0.0",
        "timestamp": "2026-04-10T10:00:00Z",
    })
}

func (h *Handler) GetDemoProcesses(c *gin.Context) {
    processes := []gin.H{
        {
            "name":          "MonetizationArtifactGeneration",
            "description":   "Transforms ingested data and business rules into monetization models",
            "inputEntities": []string{"IngestedData", "BusinessRules", "PricingParameters"},
            "outputEntities": []string{"MonetizationModel", "AnalyticsOffering", "RateCard"},
        },
        {
            "name":        "DataIngestionTrigger",
            "description": "Occurs when new market or operational data is ingested",
            "triggerType": "event",
            "frequency":   "real-time",
        },
        {
            "name":               "DynamicPricingOptimization",
            "description":        "Enables data-driven creation of monetization models",
            "optimizationTarget": "revenue",
            "algorithm":          "machine-learning",
        },
    }

    c.JSON(http.StatusOK, models.DemoResponse{
        Status:  "success",
        Message: "Demo processes retrieved successfully",
        Data:    processes,
    })
}
