package db

import (
	"app/model"
)

func (d *DataStore) GetListApplication() ([]model.Application, error) {
	var Applications []model.Application
	res := d.Db.Preload("Owner.Profile").Preload("Lead.Profile").Find(&Applications)
	if res.Error != nil {
		return nil, res.Error
	}
	return Applications, nil
}

func (d *DataStore) InsertApplication(Application *model.Application) error {
	return d.Db.Create(Application).Error
}

func (d *DataStore) DeleteApplicationByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.Application{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) UpdateApplication(Application *model.Application) error {
	err := d.Db.Save(&Application).Error
	if err != nil {
		return err
	}
	return nil
}
