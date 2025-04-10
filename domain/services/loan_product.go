package services

import (
	"context"

	"github.com/JamshedJ/loanly/domain/repository"
	"github.com/rs/zerolog"
)

type LoanProductService struct {
	Logger          zerolog.Logger
	LoanProductRepo repository.LoanProductRepositoryI
}

var _ LoanProductServiceI = (*LoanProductService)(nil)

func (s *LoanProductService) Create(ctx context.Context, in CreateLoanProductIn) (uint, error) {
	logger := s.Logger.With().Ctx(ctx).Str("handler", "CreateLoanProduct").Logger()

	buildedLP, err := LPBuilder(in)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to build loan product")
		return 0, err
	}

	lpId, err := s.LoanProductRepo.Create(ctx, buildedLP)
	if err != nil {
		logger.Error().Err(err).Msg("Failed create loan product")
		return 0, err
	}

	return lpId, nil
}
