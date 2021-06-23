package repository

import "github.com/john-victor-100/CreditTodxs/internal/app/domain"

// IProposal is interface of order repository
type IProposal interface {
	Get() domain.Proposal
	Accept(domain.Client, domain.Proposal) domain.Loan
}
