package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"loan-rest-api/api/DTO"
	"loan-rest-api/model"
)

type LoanController struct {
	DB *gorm.DB
}

func (c *LoanController) GetLoan(context *gin.Context) {
	var loanGetRequest DTO.LoanGetRequest
	context.BindUri(&loanGetRequest)
	var loan model.Loan
	c.DB.First(&loan, loanGetRequest.ID)
	if loan.ID == 0 {
		context.JSON(404, "{success:false, message: 'loan not found with this ID'}")
	} else {
		context.JSON(200, loan)
	}
}

func (c *LoanController) PostLoan(context *gin.Context) {
	var loanRequest *DTO.LoanPostRequest
	context.Bind(&loanRequest)
	loanModel := loanRequest.GetModel()
	c.DB.Create(&loanModel)
	context.JSON(200, loanModel)
}
