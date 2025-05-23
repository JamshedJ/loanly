package v1

import (
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func (api *AdminApiV1) CreateLoanProduct(c *gin.Context) {
	logger := api.Logger.With().Ctx(c).Str("handler", "CreateLoanProduct").Logger()

	in := struct {
		Name            string          `json:"name"`
		Type            string          `json:"type"`
		MinAmount       decimal.Decimal `json:"min_amount"`
		MaxAmount       decimal.Decimal `json:"max_amount"`
		Currency        string          `json:"currency"`
		MinTermDays     int             `json:"min_term_days"`
		MaxTermDays     int             `json:"max_term_days"`
		MinInterestRate decimal.Decimal `json:"min_interest_rate"`
		MaxInterestRate decimal.Decimal `json:"max_interest_rate"`
	}{}

	if err := c.ShouldBindJSON(&in); err != nil {
		logger.Error().Err(err).Msg("failed to bind json")
		handleError(c, err)
		return
	}

	err := api.Svc.LoanProduct.Create(c, services.CreateLoanProductIn{
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

	c.JSON(200, gin.H{"loan_product": "created"})
}
