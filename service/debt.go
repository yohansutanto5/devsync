package service

import (
	"app/db"
	"app/model"
)

// DebtService defines the interface for managing Debts.
type DebtService interface {
	Insert(Debt *model.Debt) error
	GetByID(id int) (*model.Debt, error)
	Update(data *model.Debt) error
	DeleteByID(id int) error
	GetList() ([]model.Debt, error)
}

type DebtServiceImpl struct {
	db *db.DataStore
}

func NewDebtService(db *db.DataStore) DebtService {
	return &DebtServiceImpl{db: db}
}

// Function Implementation
func (s *DebtServiceImpl) GetByID(id int) (*model.Debt, error) {
	// Implementation for fetching a Debt by ID from the database
	Debt := &model.Debt{}
	if err := s.db.Db.Delete(Debt).Error; err != nil {
		return nil, err
	}
	return Debt, nil
}

func (s DebtServiceImpl) GetList() ([]model.Debt, error) {
	return s.db.GetListDebt()
}

func (s *DebtServiceImpl) DeleteByID(id int) error {
	return s.db.DeleteDebtByID(id)
}

func (s *DebtServiceImpl) Update(data *model.Debt) error {
	return s.db.UpdateDebt(data)
}

func (s *DebtServiceImpl) Insert(Debt *model.Debt) error {
	return s.db.InsertDebt(Debt)
}
