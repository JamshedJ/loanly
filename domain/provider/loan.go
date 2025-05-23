package provider

import "context"

type LoanProviderI interface {
	CreateLoanProduct(ctx context.Context, in CreateLoanProductIn) error
	CreateLoan(ctx context.Context, in CreateLoanIn) error
	ActivateLoan(ctx context.Context, id string) error
	IssueLoan(ctx context.Context, id string) error
	RepayLoan(ctx context.Context, id string) error
}

type CreateLoanProductIn struct {}

type CreateLoanIn struct {}
