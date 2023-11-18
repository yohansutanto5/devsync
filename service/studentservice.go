package service

import (
	"app/db"
	"app/model"
)

// StudentService defines the interface for managing students.
type StudentService interface {
	Create(student *model.Student) error
	GetByID(id int) (*model.Student, error)
	Update(data *model.Student) error
	DeleteByID(id int) error
	New(FirstName, LastName string, id int) model.Student
	GetList() ([]model.Student, error)
}

type StudentServiceImpl struct {
	db *db.DataStore
}

func NewStudentService(db *db.DataStore) StudentService {
	return &StudentServiceImpl{db: db}
}

// Function Implementation
func (s *StudentServiceImpl) GetByID(id int) (*model.Student, error) {
	// Implementation for fetching a student by ID from the database
	student := &model.Student{}
	if err := s.db.Db.Delete(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s StudentServiceImpl) GetList() ([]model.Student, error) {
	return s.db.GetListStudent()
}

func (s *StudentServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteStudentByID(id)
}

func (s *StudentServiceImpl) Update(data *model.Student) error {
	return s.db.UpdateStudent(data)
}

func (s *StudentServiceImpl) Create(student *model.Student) error {
	return s.db.InsertStudent(student)
}

func (s *StudentServiceImpl) New(FirstName, LastName string, id int) model.Student {
	var st model.Student
	st.FirstName = FirstName
	st.LastName = LastName
	st.ID = id
	return st
}
