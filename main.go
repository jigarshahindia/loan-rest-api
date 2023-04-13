package main

import (
	"loan-rest-api/api"
	"loan-rest-api/storage"
)

/*
Requirements for Loan API

GET
	Get by LoanID/CustomerID ?
	Return list/singleLoan
POST
	Loan
 	Save in some store
*/
func main() {
	command()
}

func command() {
	storage.Connect()
	storage.Migrate()
	api.StartServer()
}
