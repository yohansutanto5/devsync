package service

import (
	"app/db"
	"app/model"
)

// UserService defines the interface for managing Users.
type UserService interface {
	Insert(User *model.User) error
	GetByID(id int) (*model.User, error)
	Update(data *model.User) error
	DeleteByID(id int) error
	New(FirstName, LastName string, id int) model.User
	GetList() ([]model.User, error)
}

type UserServiceImpl struct {
	db *db.DataStore
}

func NewUserService(db *db.DataStore) UserService {
	return &UserServiceImpl{db: db}
}

// Function Implementation
func (s *UserServiceImpl) GetByID(id int) (*model.User, error) {
	// Implementation for fetching a User by ID from the database
	User := &model.User{}
	if err := s.db.Db.Delete(User).Error; err != nil {
		return nil, err
	}
	return User, nil
}

func (s UserServiceImpl) GetList() ([]model.User, error) {
	return s.db.GetListUser()
}

func (s *UserServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteUserByID(id)
}

func (s *UserServiceImpl) Update(data *model.User) error {
	return s.db.UpdateUser(data)
}

func (s *UserServiceImpl) Insert(User *model.User) error {
	return s.db.InsertUser(User)
}

func (s *UserServiceImpl) New(FirstName, LastName string, id int) model.User {
	var st model.User
	st.FirstName = FirstName
	st.LastName = LastName
	st.ID = id
	return st
}
