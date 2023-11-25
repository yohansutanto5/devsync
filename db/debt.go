package db

import (
	"app/model"
)

func (d *DataStore) GetListDebt() ([]model.Debt, error) {
	var Debts []model.Debt
	res := d.Db.Preload("Application").Find(&Debts)
	if res.Error != nil {
		return nil, res.Error
	}
	return Debts, nil
}

func (d *DataStore) InsertDebt(Debt *model.Debt) error {
	return d.Db.Create(Debt).Error
}

func (d *DataStore) DeleteDebtByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.Debt{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) UpdateDebt(Debt *model.Debt) error {
	err := d.Db.Save(&Debt).Error
	if err != nil {
		return err
	}
	return nil
}
