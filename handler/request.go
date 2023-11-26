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

func GetRequest(c *gin.Context, Request service.RequestService) {
	// To DO handle filter and search
	result, err := Request.GetList()

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

func DeleteRequest(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var Request = service.NewRequestService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, Request.DeleteByID(id))
}
