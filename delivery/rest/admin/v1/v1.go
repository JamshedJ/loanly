package v1

import (
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type AdminApiV1 struct {
	Logger zerolog.Logger
	Svc    services.ServiceFacade
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Origin")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (api *AdminApiV1) RegisterRoutes() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(CORS())

	v1 := e.Group("admin/v1")
	{
		lp := v1.Group("/loanproduct")
		{
			lp.POST("", api.CreateLoanProduct)
		}
	}

	return e
}
