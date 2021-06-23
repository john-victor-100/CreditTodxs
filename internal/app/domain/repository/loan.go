package repository

import "github.com/john-victor-100/CreditTodxs/internal/app/domain"

// ILoan is interface of order repository
type ILoan interface {
	Get() domain.Loan
	Update(domain.Loan) // receives Loan with updated values and tries to update loan data
}
