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

func GetApplication(c *gin.Context, Application service.ApplicationService) {
	// To DO handle filter and search
	result, err := Application.GetList()

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

func InsertApplication(c *gin.Context, Application service.ApplicationService) {
	// Cast data from request
	var data model.InsertApplicationIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct Application Model with the request data
	var newApplication model.Application
	util.ConvertStruct(data, &newApplication)

	// Call create service
	err := Application.Insert(&newApplication)

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

func DeleteApplication(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var Application = service.NewApplicationService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, Application.DeleteByID(id))
}

// func UpdateApplication(c *gin.Context) {
// 	dbService := db.GetContext(c)
// 	var data model.AddApplicationIn
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var Application = service.NewApplicationService(dbService)
// 	var newApplication model.Application
// 	newApplication.FirstName = data.FirstName
// 	newApplication.LastName = data.LastName

// 	err := Application.Update(&newApplication)
// 	if err != nil {
// 		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
// 		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
// 	}
// 	c.JSON(http.StatusOK, constanta.SuccessMessage)
// }
