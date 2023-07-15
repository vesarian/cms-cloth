package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewInternalServerError(c *gin.Context, error string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   "Internal Server Error",
		"message": error,
	})
}
