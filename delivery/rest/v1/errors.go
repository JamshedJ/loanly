package v1

import (
	"errors"

	"github.com/JamshedJ/loanly/infrastructure/repository/errs"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrRecordNotFound):
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
		return
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
}
