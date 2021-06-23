package usecase

import (
	"github.com/john-victor-100/CreditTodxs/internal/app/domain"
	"github.com/john-victor-100/CreditTodxs/internal/app/domain/repository"
)

// ProposalGet is the usecase of returning a proposal
func ProposalGet(p repository.IProposal) domain.Proposal {
	return p.Get()
}

// ProposalAccept is the usecase of accepting a loan
func ProposalAccept(c repository.IClient, p repository.IProposal) domain.Loan {
	return p.Accept(c.Get(), p.Get())
}
