package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/log"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, User service.UserService) {
	// To DO handle filter and search
	result, err := User.GetList()

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

func InsertUser(c *gin.Context, User service.UserService) {
	// Cast data from request
	var data model.AddUserIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct User Model with the request data
	var newUser model.User
	newUser.FirstName = data.FirstName
	newUser.LastName = data.LastName
	newUser.Active = true
	newUser.Email = data.Email
	newUser.Username = data.Username
	newUser.ProfileID = data.Profile

	// Call create service
	err := User.Insert(&newUser)

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

func DeleteUser(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var User = service.NewUserService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, User.DeleteByID(id))
}

func UpdateUser(c *gin.Context) {
	dbService := db.GetContext(c)
	var data model.AddUserIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var User = service.NewUserService(dbService)
	var newUser model.User
	newUser.FirstName = data.FirstName
	newUser.LastName = data.LastName

	err := User.Update(&newUser)
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	}
	c.JSON(http.StatusOK, constanta.SuccessMessage)
}
