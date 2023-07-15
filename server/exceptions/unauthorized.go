package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUnauthorizedError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "Unauthorized",
		"message": message,
	})
}
