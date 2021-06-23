package repository

import "github.com/john-victor-100/CreditTodxs/internal/app/domain"

// IClient is interface of order repository
type IClient interface {
	Get() domain.Client
	Register(domain.Client)
}
