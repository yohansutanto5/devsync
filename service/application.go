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
	return s.db.InsertApplication(Application)
}