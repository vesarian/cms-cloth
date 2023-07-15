package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewNotFoundError(c *gin.Context, error string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   "Data Not Found",
		"message": error,
	})
}
