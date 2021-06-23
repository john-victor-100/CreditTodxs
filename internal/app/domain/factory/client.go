package factory

import (
	"github.com/john-victor-100/CreditTodxs/internal/app/domain"
)

// Order is the factory of domain.Order
type Client struct{}

// Generate generates domain.Client from primitives
func (of Client) Generate(
	ID string,
	Name string,
	CPF int64,
	Birthday int64,
	CreditScore int,
	Age int,
	Married bool,
	Children int,
	RealStateCollateral bool,
	RealStateCollateralValue float64,
	CarCollateral bool,
	CarCollateralValue bool,
) domain.Client {
	return domain.Client{
		ID:                       ID,
		Name:                     Name,
		CPF:                      CPF,
		Birthday:                 Birthday,
		CreditScore:              CreditScore,
		Age:                      Age,
		Married:                  Married,
		Children:                 Children,
		RealStateCollateral:      RealStateCollateral,
		RealStateCollateralValue: RealStateCollateralValue,
		CarCollateral:            CarCollateral,
		CarCollateralValue:       CarCollateralValue,
	}
}
