package main

import (
    "log"
    "go-api-service/handlers"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
    router.Use(cors.New(config))

    h := handlers.NewHandler()

    api := router.Group("/api/v1")
    {
        api.POST("/schemas/:schemaId/instances", h.CreateSchemaInstances)
        api.GET("/entities/static", h.GetStaticEntities)
        api.GET("/stats", h.GetSystemStats)
        api.GET("/health", h.GetHealthCheck)
        api.GET("/processes/demo", h.GetDemoProcesses)
    }

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Go API Service is running",
            "endpoints": []string{
                "POST /api/v1/schemas/:schemaId/instances - Create schema instances",
                "GET  /api/v1/entities/static - Get static entities",
                "GET  /api/v1/stats - Get system statistics",
                "GET  /api/v1/health - Health check",
                "GET  /api/v1/processes/demo - Get demo processes",
            },
        })
    })

    log.Println("Server starting on :3000")
    if err := router.Run(":3000"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
