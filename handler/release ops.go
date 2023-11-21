package handler

import (
	"app/constanta"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReleaseTicket(c *gin.Context, ReleaseOPS service.ReleaseOPSService) {
	// To DO handle filter and search
	result, err := ReleaseOPS.GetListTicket()

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

func TriggerBuild(c *gin.Context, ReleaseOPS service.ReleaseOPSService) {
	// To DO handle filter and search
	err := ReleaseOPS.TriggerBuild(util.ConvertToInt(c.Param("ID")))
	if err != nil {
		errorResponse := model.ErrorResponse{
			TransactionID: util.GetTransactionID(c),
			Message:       constanta.InternalServerErrorMessage,
			Code:          constanta.CodeErrorService,
			Details:       err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	} else {
		c.JSON(http.StatusOK, "result")
	}
}

func InsertReleaseTicket(c *gin.Context, ReleaseOPS service.ReleaseOPSService) {
	// Cast data from request
	var data model.InsertReleaseTicketIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct ReleaseOPS Model with the request data
	var newReleaseTicket = model.ReleaseTicket{
		AppID:      data.AppID,
		Job:        data.Job,
		VersionUAT: data.VersionUAT,
		VersionPRD: data.VersionPRD,
		Workflow:   data.Workflow,
		Status:     "READY",
	}

	// Call create service
	err := ReleaseOPS.Insert(&newReleaseTicket)

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

// func DeleteReleaseOPS(c *gin.Context) {
// 	dbService := db.GetContext(c)
// 	// To do parsing data here
// 	id := 1
// 	var ReleaseOPS = service.NewReleaseOPSService(dbService)
// 	// To DO handle filter and search
// 	c.JSON(http.StatusOK, ReleaseOPS.DeleteByID(id))
// }

// func UpdateReleaseOPS(c *gin.Context) {
// 	dbService := db.GetContext(c)
// 	var data model.AddReleaseOPSIn
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var ReleaseOPS = service.NewReleaseOPSService(dbService)
// 	var newReleaseOPS model.ReleaseOPS
// 	newReleaseOPS.FirstName = data.FirstName
// 	newReleaseOPS.LastName = data.LastName

// 	err := ReleaseOPS.Update(&newReleaseOPS)
// 	if err != nil {
// 		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
// 		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
// 	}
// 	c.JSON(http.StatusOK, constanta.SuccessMessage)
// }
