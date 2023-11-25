package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDebt(c *gin.Context, Debt service.DebtService) {
	// To DO handle filter and search
	result, err := Debt.GetList()

	if err != nil {
		errorResponse := model.ErrorResponse{
			TransactionID: util.GetTransactionID(c),
			Message:       constanta.InternalServerErrorMessage,
			Code:          constanta.CodeErrorService,
			Details:       err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func InsertDebt(c *gin.Context, Debt service.DebtService) {
	// Cast data from request
	var data model.InsertDebtIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct Debt Model with the request data
	var newDebt model.Debt
	util.ConvertStruct(data, &newDebt)

	// Call create service
	err := Debt.Insert(&newDebt)

	// Construct Response
	if err != nil {
		errorResponse := model.ErrorResponse{
			TransactionID: util.GetTransactionID(c),
			Message:       constanta.InternalServerErrorMessage,
			Code:          constanta.CodeErrorService,
			Details:       err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	} else {
		c.JSON(http.StatusOK, constanta.SuccessMessage)
	}
}

func DeleteDebt(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var Debt = service.NewDebtService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, Debt.DeleteByID(id))
}

// func UpdateDebt(c *gin.Context) {
// 	dbService := db.GetContext(c)
// 	var data model.AddDebtIn
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var Debt = service.NewDebtService(dbService)
// 	var newDebt model.Debt
// 	newDebt.FirstName = data.FirstName
// 	newDebt.LastName = data.LastName

// 	err := Debt.Update(&newDebt)
// 	if err != nil {
// 		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
// 		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
// 	}
// 	c.JSON(http.StatusOK, constanta.SuccessMessage)
// }
