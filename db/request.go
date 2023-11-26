package db

import (
	"app/model"
)

func (d *DataStore) GetListRequest() ([]model.Request, error) {
	var Requests []model.Request
	res := d.Db.Find(&Requests)
	if res.Error != nil {
		return nil, res.Error
	}
	return Requests, nil
}

func (d *DataStore) InsertRequest(Request *model.Request) error {
	return d.Db.Create(Request).Error
}

func (d *DataStore) DeleteRequestByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.Request{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
