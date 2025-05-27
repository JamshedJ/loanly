package provider

import "context"

type LoanProviderI interface {
	// LoanProduct
	CreateLoanProduct(ctx context.Context, in CreateLoanProductIn) error
	// Loan
	CreateLoan(ctx context.Context, in CreateLoanIn) error
	ActivateLoan(ctx context.Context, id string) error
	IssueLoan(ctx context.Context, id string) error
	RepayLoan(ctx context.Context, id string) error
}

type CreateLoanProductIn struct {}

type CreateLoanIn struct {}
