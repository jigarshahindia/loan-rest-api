package DTO

import (
	"loan-rest-api/model"
)

type LoanPostRequest struct {
	CustomerID int    `json:"CustomerID"`
	Amount     int    `json:"Amount"`
	Status     string `json:"Status"`
}

type LoanGetRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (request *LoanPostRequest) GetModel() model.Loan {
	return model.Loan{
		CustomerID: request.CustomerID,
		Amount:     request.Amount,
		Status:     request.Status,
	}
}
