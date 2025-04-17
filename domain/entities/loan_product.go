package entities

import "github.com/shopspring/decimal"

type LoanProduct struct {
	ID                   string               `json:"id"`
	Name                 string               `json:"name"`
	Type                 LoanProductType      `json:"type"`
	MinAmount            decimal.Decimal      `json:"min_amount"`
	MaxAmount            decimal.Decimal      `json:"max_amount"`
	Currency             Currency             `json:"currency"`
	MinTermDays          int                  `json:"min_term_days"`
	MaxTermDays          int                  `json:"max_term_days"`
	MinInterestRate      decimal.Decimal      `json:"interest_rate"`         // Минимальная процентная ставка
	MaxInterestRate      decimal.Decimal      `json:"max_interest_rate"`     // Максимальная процентная ставка
	LateFee              decimal.Decimal      `json:"late_fee"`              // Штраф за просрочку в процентах
	InstallmentFrequency InstallmentFrequency `json:"installment_frequency"` // Частота платежей (еженедельно, ежемесячно и т.д.)
	SinglePayment        bool                 `json:"single_payment"`        // Погашение одной выплатой или нет
	Installments         []int                `json:"installments"`          // Количество платежей
}

type LoanProductType string

const (
	LoanProductTypePayday      LoanProductType = "payday"      // Кредит до зарплаты
	LoanProductTypeInstallment LoanProductType = "installment" // Классический кредит
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	RUB Currency = "RUB"
	TJS Currency = "TJS"
)

type InstallmentFrequency string

const (
	InstallmentFrequencyWeekly  InstallmentFrequency = "weekly"
	InstallmentFrequencyMonthly InstallmentFrequency = "monthly"
)
