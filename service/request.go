package service

import (
	"app/db"
	"app/model"
)

// RequestService defines the interface for managing Requests.
type RequestService interface {
	Insert(Request *model.Request) error
	GetByID(id int) (*model.Request, error)
	DeleteByID(id int) error
	GetList() ([]model.Request, error)
}

type RequestServiceImpl struct {
	db *db.DataStore
}

func NewRequestService(db *db.DataStore) RequestService {
	return &RequestServiceImpl{db: db}
}

// Function Implementation
func (s *RequestServiceImpl) GetByID(id int) (*model.Request, error) {
	// Implementation for fetching a Request by ID from the database
	Request := &model.Request{}
	if err := s.db.Db.Delete(Request).Error; err != nil {
		return nil, err
	}
	return Request, nil
}

func (s RequestServiceImpl) GetList() ([]model.Request, error) {
	return s.db.GetListRequest()
}

func (s *RequestServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteRequestByID(id)
}

func (s *RequestServiceImpl) Insert(Request *model.Request) error {
	return s.db.InsertRequest(Request)
}
