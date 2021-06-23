package model

// Loan has all the info pertaining to a loan of a client
type Loan struct {
	ID              string `gorm:"column:card_id;type:uuid"`
	ClientID        string `gorm:"type:uuid"`
	Client          Client `gorm:"foreignKey:ClientID;references:ID"`
	Interest        float64
	Installments    int     //Number of payments
	Installment     float64 // Value per installment
	InstallmentNext int     // Next installment index e.g if 12 total, next paymeny will be third so 3 will be the InstallmentNext value
	AmountTotal     float64
	AmountPaid      float64
}
