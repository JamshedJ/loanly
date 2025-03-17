package services

import "github.com/JamshedJ/loanly/domain/provider"

type ServiceFacade struct {
	LoanProvider provider.LoanProviderI
}