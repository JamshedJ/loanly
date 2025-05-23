package services

import (
	"context"
)

type LoanProductServiceI interface {
	Create(ctx context.Context, in CreateLoanProductIn) error
}
