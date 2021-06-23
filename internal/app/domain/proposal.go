package domain

// Proposal is a loan offer that is offered to a given client
type Proposal struct {
	ID           string
	Interest     float64
	Installments int     //Number of payments
	Installment  float64 // Value per installment
	Amount       float64
	FinalDebt    float64 // Total amount after interest
}
