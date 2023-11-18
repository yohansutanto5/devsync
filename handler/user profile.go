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

func GetUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	// To DO handle filter and search
	result, err := UserProfile.GetList()

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

func DeleteUserProfile(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var UserProfile = service.NewUserProfileService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, UserProfile.DeleteByID(id))
}

func InsertUserProfile(c *gin.Context, UserProfile service.UserProfileService) {
	// Cast data from request
	var data model.AddUserProfileIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct UserProfile Model with the request data
	var newUserProfile model.UserProfile
	newUserProfile.Name = data.Name

	// Call create service
	err := UserProfile.Insert(&newUserProfile)

	// Construct Response
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage+util.GetTransactionIDString(c))
	} else {
		c.JSON(http.StatusOK, constanta.SuccessMessage)
	}
}

func UpdateUserProfile(c *gin.Context) {
	dbService := db.GetContext(c)
	var data model.AddUserProfileIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var UserProfile = service.NewUserProfileService(dbService)
	var newUserProfile model.UserProfile
	newUserProfile.Name = data.Name

	err := UserProfile.Update(&newUserProfile)
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	}
	c.JSON(http.StatusOK, constanta.SuccessMessage)
}
