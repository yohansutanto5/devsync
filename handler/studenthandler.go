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

func GetStudent(c *gin.Context, student service.StudentService) {
	// To DO handle filter and search
	result, err := student.GetList()

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

func DeleteStudent(c *gin.Context) {
	dbService := db.GetContext(c)
	// To do parsing data here
	id := 1
	var student = service.NewStudentService(dbService)
	// To DO handle filter and search
	c.JSON(http.StatusOK, student.DeleteByID(id))
}

func AddStudent(c *gin.Context, student service.StudentService) {
	// Cast data from request
	var data model.AddStudentIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct Student Model with the request data
	var newStudent model.Student
	newStudent.FirstName = data.FirstName
	newStudent.LastName = data.LastName

	// Call create service
	err := student.Create(&newStudent)

	// Construct Response
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	} else {
		c.JSON(http.StatusOK, constanta.SuccessMessage)
	}
}

func UpdateStudent(c *gin.Context) {
	dbService := db.GetContext(c)
	var data model.AddStudentIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var student = service.NewStudentService(dbService)
	var newStudent model.Student
	newStudent.FirstName = data.FirstName
	newStudent.LastName = data.LastName

	err := student.Update(&newStudent)
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), constanta.InternalServerErrorCode, nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	}
	c.JSON(http.StatusOK, constanta.SuccessMessage)
}
