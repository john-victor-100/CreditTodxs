package repository

import (
	"github.com/john-victor-100/CreditTodxs/internal/app/adapter/postgresql"
	"github.com/john-victor-100/CreditTodxs/internal/app/adapter/postgresql/model"
	"github.com/john-victor-100/CreditTodxs/internal/app/domain"
	"github.com/john-victor-100/CreditTodxs/internal/app/domain/factory"
)

// Client is the repository of domain.Order
type Client struct{}

// Get gets client
func (C Client) Get() domain.Client {
	db := postgresql.Connection()
	var client model.Client
	// Order has Person/Payment relation and Payment has Card relation which has CardBrand relation.
	result := db.First(&client, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	var c factory.Client
	return c.Generate(
		client.ID,
		client.Name,
		client.CPF,
		client.Birthday,
		client.CreditScore,
		client.Age,
		client.Married,
		client.Children,
		client.RealStateCollateral,
		client.RealStateCollateralValue,
		client.CarCollateral,
		client.CarCollateralValue,
	)
}

func (C Client) Register() {

}
