package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondWithError sends a JSON response with an error message and status code
func RespondWithError(c *gin.Context, code int, message string) {
	Logger.Error(message)
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// RespondWithSuccess sends a JSON response with data and a status code 200
func RespondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
