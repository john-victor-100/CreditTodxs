package usecase

import (
	"github.com/john-victor-100/CreditTodxs/internal/app/domain"
	"github.com/john-victor-100/CreditTodxs/internal/app/domain/repository"
)

// Client is the usecase of getting client
func ClientGet(c repository.IClient) domain.Client {
	return c.Get()
}

// Client is the usecase of registering client
func ClientRegister(c repository.IClient) {
	c.Register(c.Get())
}
