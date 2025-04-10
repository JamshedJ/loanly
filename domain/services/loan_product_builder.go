package services

import (
	"time"

	"github.com/JamshedJ/loanly/domain/entities"
	"github.com/JamshedJ/loanly/domain/errs"
	"github.com/shopspring/decimal"
)

// Function to check if a year is a leap year
func isLeapYear(year int) bool {
	// A year is a leap year if it is divisible by 4, but not divisible by 100, unless it is also divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// This function checks if a given year is a leap year (высокосный год)
func calculateTermWithLeapYear(minTermDays, maxTermDays int) (int, int) {
	// First, we need define the start date
	startDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Calculate the minimum and maximum end dates based on the start date
	minEndDate := startDate.AddDate(0, 0, minTermDays)
	maxEndDate := startDate.AddDate(0, 0, maxTermDays)

	// Check if we have a leap year in the range of minEndDate and maxEndDate
	for date := minEndDate; date.Before(maxEndDate); date = date.AddDate(0, 0, 1) {
		if isLeapYear(date.Year()) {
			// If we find a leap year, we need to add 1 day to both minTermDays and maxTermDays
			minTermDays++
			maxTermDays++
			break
		}
	}

	return minTermDays, maxTermDays
}

func LPBuilder(in CreateLoanProductIn) (entities.LoanProduct, error) {
	if err := in.Validate(); err != nil {
		return entities.LoanProduct{}, errs.ErrValidationFailed
	}

	lpTypes := []entities.LoanProductType{
		entities.LoanProductTypeBNPL,
		entities.LoanProductTypePayday,
		entities.LoanProductTypeInstallment,
	}

	lp := entities.LoanProduct{}
	
	minTermD, maxTermD := calculateTermWithLeapYear(in.MinTermDays, in.MaxTermDays)

	for _, lptype := range lpTypes {
		switch lptype {
		case entities.LoanProductTypeBNPL:
			lp = entities.LoanProduct{
				Name:            in.Name,
				Type:            lptype,
				MinAmount:       in.MinAmount,
				MaxAmount:       in.MaxAmount,
				Currency:        entities.Currency(in.Currency),
				MinTermDays:     minTermD,
				MaxTermDays:     maxTermD,
				MinInterestRate: in.MinInterestRate,
				MaxInterestRate: in.MaxInterestRate,
				// TODO: Написать формулу расчета LateFee
				LateFee:              decimal.NewFromFloat(0.5), // 0.5 TJS в день
				InstallmentFrequency: entities.InstallmentFrequencyMonthly,
				SinglePayment:        false,
				Installments:         []int{1, 3, 6, 12, 24},
			}
			return lp, nil
		case entities.LoanProductTypePayday:
			lp = entities.LoanProduct{
				Name:                 in.Name,
				Type:                 lptype,
				MinAmount:            in.MinAmount,
				MaxAmount:            in.MaxAmount,
				Currency:             entities.Currency(in.Currency),
				MinTermDays:          minTermD,
				MaxTermDays:          maxTermD,
				MinInterestRate:      in.MinInterestRate,
				MaxInterestRate:      in.MaxInterestRate,
				LateFee:              decimal.NewFromFloat(0.50), // 0.5 TJS в день (каждый день будет начисляться штраф 0.5 TJS до полного погашения)
				InstallmentFrequency: entities.InstallmentFrequencyMonthly,
				SinglePayment:        true,
				Installments:         []int{1},
			}
		case entities.LoanProductTypeInstallment:
			lp = entities.LoanProduct{
				Name:                 in.Name,
				Type:                 lptype,
				MinAmount:            in.MinAmount,
				MaxAmount:            in.MaxAmount,
				Currency:             entities.Currency(in.Currency),
				MinTermDays:          minTermD,
				MaxTermDays:          maxTermD,
				MinInterestRate:      in.MinInterestRate,
				MaxInterestRate:      in.MaxInterestRate,
				LateFee:              decimal.NewFromFloat(0.1),
				InstallmentFrequency: entities.InstallmentFrequencyMonthly,
				SinglePayment:        false,
				Installments:         []int{1, 3, 6, 12, 24},
			}
		}
	}
	return lp, nil
}
