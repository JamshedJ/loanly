package v1

import (
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type LoanApiV1 struct {
	Logger zerolog.Logger
	Svc    services.ServiceFacade
}

func (api *LoanApiV1) RegisterRoutes(e *gin.Engine) error {
	e.Use(gin.Recovery())

	v1 := e.Group("v1")
	{
		loan := v1.Group("/loan")
		{
			loan.POST("")
		}
	}
	return nil
}
