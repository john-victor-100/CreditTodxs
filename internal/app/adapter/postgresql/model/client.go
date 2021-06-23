package model

// Client is the structure that holds the client and all its data
type Client struct {
	ID                       string `gorm:"column:order_id;type:uuid"`
	Name                     string
	CPF                      int64
	Birthday                 int64 // Epoch time
	CreditScore              int
	Age                      int
	Married                  bool
	Children                 int  //How many children a person has
	RealStateCollateral      bool // If client offers real state as collateral
	RealStateCollateralValue float64
	CarCollateral            bool // If client offers vehicle(s) as collateral
	CarCollateralValue       bool
}
