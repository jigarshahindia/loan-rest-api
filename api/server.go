package api

import (
	"github.com/gin-gonic/gin"
	"loan-rest-api/controller"
	"loan-rest-api/storage"
)

func StartServer() {
	r, _ := router()
	r.Run()
}

func router() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.BasicAuth(gin.Accounts{
		"test": "test",
	}))

	r.GET("/ping", controller.Ping)
	loanController := controller.LoanController{DB: storage.LOAN_DB}

	r.GET("/loan/:id", loanController.GetLoan)
	r.POST("/loan", loanController.PostLoan)
	return r, nil
}
