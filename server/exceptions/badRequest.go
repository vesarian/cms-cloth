package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBadRequestError(c *gin.Context, error string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   "Bad Request",
		"message": error,
	})
}
