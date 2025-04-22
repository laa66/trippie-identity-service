package ctx

import "github.com/gin-gonic/gin"

type CorrelationData map[string]any

type AppContext struct {
	Context         *gin.Context
	CorrelationData CorrelationData
}

func (c *AppContext) GetCorrelationData() CorrelationData {
	return c.CorrelationData
}
