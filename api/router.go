package api

import "github.com/gin-gonic/gin"

// Register endpoints
func SetupRouter() *gin.Engine {
    r := gin.Default()

	api := NewApi()
	api.RegisterIdentityEndpoints(r.Group("/identity"))

    return r
}