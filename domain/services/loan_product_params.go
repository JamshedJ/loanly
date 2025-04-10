package services

import (
	"fmt"

	"github.com/JamshedJ/loanly/domain/entities"
	"github.com/shopspring/decimal"
)

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

// TODO: В комментариях показать в виде таблицы каждого типа продукта и его параметры

func (i *CreateLoanProductIn) Validate() error {
	switch i.Type {
	case string(entities.LoanProductTypeBNPL):
		if i.MinAmount.Cmp(decimal.NewFromFloat(500)) < 0 {
			return fmt.Errorf("min_amount for BNPL must be greater than 500")
		}
		if i.MaxAmount.Cmp(decimal.NewFromFloat(25000)) > 0 {
			return fmt.Errorf("max_amount for BNPL must be less than 25000")
		}
		if i.MinTermDays < 10 {
			return fmt.Errorf("min_term_days for BNPL must be greater than 10")
		}
		if i.MaxTermDays > 730 {
			return fmt.Errorf("max_term_days for BNPL must be less than 730")
		}
		if i.MinInterestRate.Cmp(decimal.NewFromFloat(4)) < 0 {
			return fmt.Errorf("min_interest_rate for BNPL must be greater than 4%%")
		}
		if i.MaxInterestRate.Cmp(decimal.NewFromFloat(37)) > 0 {
			return fmt.Errorf("max_interest_rate for BNPL must be less than 20%%")
		}
	case string(entities.LoanProductTypePayday):
		if i.MinAmount.Cmp(decimal.NewFromFloat(100)) < 0 {
			return fmt.Errorf("min_amount for payday cannot be less than 100")
		}
		if i.MaxAmount.Cmp(decimal.NewFromFloat(10000)) > 0 {
			return fmt.Errorf("max_amount for payday cannot be greater than 10000")
		}
		if i.MinTermDays < 7 {
			return fmt.Errorf("min_term_days for payday cannot be less than 7")
		}
		if i.MaxTermDays > 30 {
			return fmt.Errorf("max_term_days for payday cannot be greater than 30")
		}
		if i.MinInterestRate.Cmp(decimal.NewFromFloat(4)) < 0 {
			return fmt.Errorf("min_interest_rate for payday cannot be less than 4%%")
		}
		if i.MaxInterestRate.Cmp(decimal.NewFromFloat(4)) > 0 {
			return fmt.Errorf("max_interest_rate for payday cannot be greater than 4%%")
		}
	case string(entities.LoanProductTypeInstallment):
		if i.MinAmount.Cmp(decimal.NewFromFloat(1000)) < 0 {
			return fmt.Errorf("min_amount for installment cannot be less than 1000")
		}
		if i.MaxAmount.Cmp(decimal.NewFromFloat(30000)) > 0 {
			return fmt.Errorf("max_amount for installment cannot be greater than 30000")
		}
		if i.MinTermDays < 30 {
			return fmt.Errorf("min_term_days for installment cannot be less than 30")
		}
		if i.MaxTermDays > 365 {
			return fmt.Errorf("max_term_days for installment cannot be greater than 365")
		}
		if i.MinInterestRate.Cmp(decimal.NewFromFloat(4)) < 0 {
			return fmt.Errorf("min_interest_rate for installment cannot be less than 4%%")
		}
		if i.MaxInterestRate.Cmp(decimal.NewFromFloat(30)) > 0 {
			return fmt.Errorf("max_interest_rate for installment cannot be greater than 30%%")
		}
	default:
		return fmt.Errorf("invalid loan product type: %s", i.Type)
	}
	return nil
}
