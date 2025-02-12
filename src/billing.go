package billing

import (
	"errors"
	"math"
	"time"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "PENDING"
	PaymentStatusPaid    PaymentStatus = "PAID"
)

// Payment represents a single payment in the loan schedule
type Payment struct {
	WeekNumber  int
	Amount      float64
	DueDate     time.Time
	Status      PaymentStatus
	PaymentDate *time.Time
}

// LoanBillingSystem manages loan payments and schedules
type LoanBillingSystem struct {
	LoanID             int
	Principal          float64
	Weeks              int
	AnnualInterestRate float64
	Payments           []Payment
}

// NewLoanBillingSystem creates a new loan billing system with initialized payment schedule
func NewLoanBillingSystem(loanID int, principal float64, weeks int, annualInterestRate float64) *LoanBillingSystem {
	lbs := &LoanBillingSystem{
		LoanID:             loanID,
		Principal:          principal,
		Weeks:              weeks,
		AnnualInterestRate: annualInterestRate,
	}
	lbs.initializeSchedule()
	return lbs
}

// initializeSchedule creates the initial payment schedule
func (lbs *LoanBillingSystem) initializeSchedule() {
	totalInterest := lbs.Principal * (lbs.AnnualInterestRate / 100)
	totalAmount := lbs.Principal + totalInterest
	weeklyPayment := math.Round(totalAmount / float64(lbs.Weeks))

	startDate := time.Now()
	lbs.Payments = make([]Payment, lbs.Weeks)

	for week := 1; week <= lbs.Weeks; week++ {
		dueDate := startDate.AddDate(0, 0, week*7)
		lbs.Payments[week-1] = Payment{
			WeekNumber:  week,
			Amount:      weeklyPayment,
			DueDate:     dueDate,
			Status:      PaymentStatusPending,
			PaymentDate: nil,
		}
	}
}

// GetOutstanding returns the current outstanding amount on the loan
func (lbs *LoanBillingSystem) GetOutstanding() float64 {
	var outstanding float64
	for _, payment := range lbs.Payments {
		if payment.Status == PaymentStatusPending {
			outstanding += payment.Amount
		}
	}
	return outstanding
}

// IsDelinquent checks if the loan is delinquent (2 or more consecutive missed payments)
func (lbs *LoanBillingSystem) IsDelinquent() bool {
	consecutiveMissed := 0
	currentDate := time.Now()

	for _, payment := range lbs.Payments {
		if payment.Status == PaymentStatusPending && payment.DueDate.Before(currentDate) {
			consecutiveMissed++
			if consecutiveMissed >= 2 {
				return true
			}
		} else {
			consecutiveMissed = 0
		}
	}
	return false
}

// MakePayment processes a payment for the loan
func (lbs *LoanBillingSystem) MakePayment(amount float64) error {
	const epsilon = 0.01
	currentDate := time.Now()

	for i := range lbs.Payments {
		if lbs.Payments[i].Status == PaymentStatusPending &&
			math.Abs(lbs.Payments[i].Amount-amount) < epsilon {
			lbs.Payments[i].Status = PaymentStatusPaid
			lbs.Payments[i].PaymentDate = &currentDate
			return nil
		}
	}
	return errors.New("no matching pending payment found for the given amount")
}

// PaymentSchedule represents the payment schedule details
type PaymentSchedule struct {
	Week        int        `json:"week"`
	Amount      float64    `json:"amount"`
	DueDate     time.Time  `json:"due_date"`
	Status      string     `json:"status"`
	PaymentDate *time.Time `json:"payment_date,omitempty"`
}

// GetPaymentSchedule returns the full payment schedule with status
func (lbs *LoanBillingSystem) GetPaymentSchedule() []PaymentSchedule {
	schedule := make([]PaymentSchedule, len(lbs.Payments))

	for i, payment := range lbs.Payments {
		schedule[i] = PaymentSchedule{
			Week:        payment.WeekNumber,
			Amount:      payment.Amount,
			DueDate:     payment.DueDate,
			Status:      string(payment.Status),
			PaymentDate: payment.PaymentDate,
		}
	}

	return schedule
}
