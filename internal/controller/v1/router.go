package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AttachRouter creates a new gin router
func AttachRouter(h *gin.Engine) {
	v1 := h.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
