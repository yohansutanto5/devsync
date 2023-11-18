package student_test

import (
	"app/model"
	"fmt"
	"testing"
)

func TestMigrationTemplate(t *testing.T) {
	err := database.Db.AutoMigrate(&model.Student{})
	if err != nil {
		fmt.Println(err.Error())
		t.Failed()
	}
}

func TestCreateTemplate(t *testing.T) {
	var student model.Student
	student.FirstName = "halo3"
	student.LastName = "babi3"

	studenService.Create(&student)
}

func TestGetListTemplate(t *testing.T) {
	students, _ := studenService.GetList()
	if len(students) < 2 {
		t.FailNow()
	}
}
