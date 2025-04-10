package v1

import (
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func (api *ApiV1) CreateLoanProduct(c *gin.Context) {
	logger := api.Logger.With().Ctx(c).Str("handler", "CreateLoanProduct").Logger()

	in := struct {
		Name            string          `json:"name" binding:"required"`
		Type            string          `json:"type" binding:"required"`
		MinAmount       decimal.Decimal `json:"min_amount" binding:"required" validate:"gt=0,ltfield=MaxAmount"`
		MaxAmount       decimal.Decimal `json:"max_amount" binding:"required" validate:"gt=0,gtfield=MinAmount"`
		Currency        string          `json:"currency" binding:"required" validate:"eq=TJS"`
		MinTermDays     int             `json:"min_term_days" binding:"required" validate:"gt=0,ltfield=MaxTermDays"`
		MaxTermDays     int             `json:"max_term_days" binding:"required" validate:"gt=0,gtfield=MinTermDays"`
		MinInterestRate decimal.Decimal `json:"min_interest_rate" binding:"required" validate:"gt=0,ltfield=MaxInterestRate"`
		MaxInterestRate decimal.Decimal `json:"max_interest_rate" binding:"required" validate:"gt=0,gtfield=MinInterestRate"`
	}{}

	if err := c.ShouldBindJSON(&in); err != nil {
		logger.Error().Err(err).Msg("failed to bind json")
		handleError(c, err)
		return
	}

	lpId, err := api.Svc.LoanProduct.Create(c, services.CreateLoanProductIn{
		Name:            in.Name,
		Type:            in.Type,
		MinAmount:       in.MinAmount,
		MaxAmount:       in.MaxAmount,
		Currency:        in.Currency,
		MinTermDays:     in.MinTermDays,
		MaxTermDays:     in.MaxTermDays,
		MinInterestRate: in.MinInterestRate,
		MaxInterestRate: in.MaxInterestRate,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create loan product")
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"loan_product_id": lpId})
}
