package services

import (
	"context"

	"github.com/JamshedJ/loanly/domain/provider"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

type LoanProductService struct {
	Logger       zerolog.Logger
	LoanProvider provider.LoanProviderI
}

var _ LoanProductServiceI = (*LoanProductService)(nil)

type CreateLoanProductIn struct {
	Name            string
	Type            string
	MinAmount       decimal.Decimal
	MaxAmount       decimal.Decimal
	Currency        string
	MinTermDays     int
	MaxTermDays     int
	MinInterestRate decimal.Decimal
	MaxInterestRate decimal.Decimal
}

func (i *CreateLoanProductIn) Validate() error {
	// TODO: Валидация параметров
	return nil
}

func (s *LoanProductService) Create(ctx context.Context, in CreateLoanProductIn) error {
	logger := s.Logger.With().Ctx(ctx).Str("handler", "CreateLoanProduct").Logger()

	if err := in.Validate(); err != nil {
		logger.Error().Err(err).Msg("Failed to validate loan product params")
		return err
	}

	err := s.LoanProvider.CreateLoanProduct(ctx, provider.CreateLoanProductIn{})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create loan product")
		return err
	}

	return nil
}
