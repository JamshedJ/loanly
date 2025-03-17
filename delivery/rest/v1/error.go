package v1

import "github.com/gin-gonic/gin"

func handleErrorWithCode(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}
