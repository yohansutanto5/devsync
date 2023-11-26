package service

import (
	"app/db"
	"app/model"
)

// ApplicationService defines the interface for managing Applications.
type ApplicationService interface {
	Insert(Application *model.Application) error
	GetByID(id int) (*model.Application, error)
	Update(data *model.Application) error
	DeleteByID(id int) error
	GetList() ([]model.Application, error)
}

type ApplicationServiceImpl struct {
	db *db.DataStore
}

func NewApplicationService(db *db.DataStore) ApplicationService {
	return &ApplicationServiceImpl{db: db}
}

// Function Implementation
func (s *ApplicationServiceImpl) GetByID(id int) (*model.Application, error) {
	// Implementation for fetching a Application by ID from the database
	Application := &model.Application{}
	if err := s.db.Db.Delete(Application).Error; err != nil {
		return nil, err
	}
	return Application, nil
}

func (s ApplicationServiceImpl) GetList() ([]model.Application, error) {
	return s.db.GetListApplication()
}

func (s *ApplicationServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteApplicationByID(id)
}

func (s *ApplicationServiceImpl) Update(data *model.Application) error {
	return s.db.UpdateApplication(data)
}

func (s *ApplicationServiceImpl) Insert(Application *model.Application) error {
	// Status should be deactivated before approved by manager
	Application.Active = false
	err := s.db.InsertApplication(Application)
	if err != nil {
		return err
	}
	// Send Activation to Email and input it as task in the user dashboard
	// The activation is either by Approve or Reject the request. Requestor also able to cancel the request
	return nil
}

func (s *ApplicationServiceImpl) Activation(appID string, action string) error {

	app, err := s.db.GetApplicationByID(appID)
	if err != nil {
		return err
	}
	// Handle rejected or approved
	if action == "APPROVED" {
		// Create jenkins Dir
		// Create Bitbucket
		app.Active = true
		app.JenkinsDir = "asdasd"
	} else {
		app.Active = false
	}

	return s.db.UpdateApplication(&app)
}
