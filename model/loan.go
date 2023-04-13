package model

import "gorm.io/gorm"

/*
Loan Model
*/
type Loan struct {
	gorm.Model
	CustomerID int
	Amount     int
	Status     string
}
