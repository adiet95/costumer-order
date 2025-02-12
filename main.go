package main

import (
	"fmt"
	billing "technical-test/src"
)

func main() {
	// Create a new loan
	loan := billing.NewLoanBillingSystem(
		100,       // loan ID
		5_000_000, // principal (5,000,000 Rp)
		50,        // weeks
		10,        // annual interest rate
	)

	// Print initial status
	fmt.Printf("Initial Outstanding Amount: %.2f\n", loan.GetOutstanding())
	fmt.Printf("Is Delinquent: %v\n", loan.IsDelinquent())

	// Make some payments
	weeklyPayment := 110000.0
	err := loan.MakePayment(weeklyPayment) // Week 1 payment
	if err != nil {
		fmt.Printf("Payment error: %v\n", err)
	}

	err = loan.MakePayment(weeklyPayment) // Week 2 payment
	if err != nil {
		fmt.Printf("Payment error: %v\n", err)
	}

	// Skip two payments to become delinquent
	// Week 3 and 4 payments are missed

	fmt.Printf("\nAfter some payments and missed payments:\n")
	fmt.Printf("Outstanding Amount: %.2f\n", loan.GetOutstanding())
	fmt.Printf("Is Delinquent: %v\n", loan.IsDelinquent())
}
