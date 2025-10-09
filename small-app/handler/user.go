package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	// Check if the size of body is more than 5KB
	if c.Request.ContentLength > 5*1024 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request body too large. Limit is 5KB"})
		return
	}

}
