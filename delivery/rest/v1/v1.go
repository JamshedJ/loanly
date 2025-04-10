package v1

import (
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ApiV1 struct {
	Logger zerolog.Logger
	Svc    services.ServiceFacade
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (api *ApiV1) RegisterRoutes() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(CORS())

	v1 := e.Group("/v1")
	{
		lp := v1.Group("/loanproduct")
		{
			lp.POST("", api.CreateLoanProduct)
		}
	}

	return e
}
