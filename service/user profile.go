package service

import (
	"app/db"
	"app/model"
)

// UserProfileService defines the interface for managing UserProfiles.
type UserProfileService interface {
	Insert(UserProfile *model.UserProfile) error
	GetByID(id int) (*model.UserProfile, error)
	Update(data *model.UserProfile) error
	DeleteByID(id int) error
	New(FirstName, LastName string, id int) model.UserProfile
	GetList() ([]model.UserProfile, error)
}

type UserProfileServiceImpl struct {
	db *db.DataStore
}

func NewUserProfileService(db *db.DataStore) UserProfileService {
	return &UserProfileServiceImpl{db: db}
}

// Function Implementation
func (s *UserProfileServiceImpl) GetByID(id int) (*model.UserProfile, error) {
	// Implementation for fetching a UserProfile by ID from the database
	UserProfile := &model.UserProfile{}
	if err := s.db.Db.Delete(UserProfile).Error; err != nil {
		return nil, err
	}
	return UserProfile, nil
}

func (s UserProfileServiceImpl) GetList() ([]model.UserProfile, error) {
	return s.db.GetListUserProfile()
}

func (s *UserProfileServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteUserProfileByID(id)
}

func (s *UserProfileServiceImpl) Update(data *model.UserProfile) error {
	return s.db.UpdateUserProfile(data)
}

func (s *UserProfileServiceImpl) Insert(UserProfile *model.UserProfile) error {
	return s.db.InsertUserProfile(UserProfile)
}

func (s *UserProfileServiceImpl) New(FirstName, LastName string, id int) model.UserProfile {
	var st model.UserProfile
	st.Name = FirstName
	st.ID = id
	return st
}
