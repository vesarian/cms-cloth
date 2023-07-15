package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUnauthenticatedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "Unauthenticated",
		"message": "Unauthenticated",
	})
}
