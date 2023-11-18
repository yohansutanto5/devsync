package application_test

import (
	"app/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockStudentService struct {
	mock.Mock
}

func (m *MockStudentService) GetByID(id int) (*model.Student, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Student), nil
}

func TestHandlerStudent(t *testing.T) {
	expected := &model.Student{
		FirstName: "asd",
		LastName:  "lolol",
		ID:        123,
	}
	mockStudentService := new(MockStudentService)
	mockStudentService.On("GetByID", 123).Return(expected, nil)

	a, _ := mockStudentService.GetByID(123)

	if a.FirstName == "asd" {
		t.FailNow()
	}
}
