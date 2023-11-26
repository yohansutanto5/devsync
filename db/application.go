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

func (d *DataStore) GetApplicationByID(id string) (res model.Application, err error) {
	err = d.Db.Where("id = ?", id).First(&res).Error
	return
}

func (d *DataStore) InsertApplication(app *model.Application, req *model.Request) error {
	tx := d.Db.Begin()
	if err := tx.Create(app).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(req).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
