package repository

import (
	"context"

	"github.com/JamshedJ/loanly/domain/entities"
)

type LoanProductRepositoryI interface {
	Create(context.Context, entities.LoanProduct) (uint, error)
}