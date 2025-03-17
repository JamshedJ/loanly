package provider

import "context"

type LoanProviderI interface {
	CreateLoan(ctx context.Context, in CreateLoanIn) error
	ActivateLoan(ctx context.Context, id string) error
	IssueLoan(ctx context.Context, id string) error
	RepayLoan(ctx context.Context, id string) error
}

type CreateLoanIn struct {}
