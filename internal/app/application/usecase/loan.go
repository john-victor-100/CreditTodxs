package usecase

import (
	"github.com/john-victor-100/CreditTodxs/internal/app/domain"
	"github.com/john-victor-100/CreditTodxs/internal/app/domain/repository"
)

// LoanGet is the usecase of returning a loaan
func LoanGet(l repository.ILoan) domain.Loan {
	return l.Get()
}

// LoanUpdate is the usecase of updating a loan
func LoanUpdate(l repository.ILoan) {
	l.Update(l.Get())
}
