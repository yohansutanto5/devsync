package db

import (
	"app/model"
)

func (d *DataStore) GetListUser() ([]model.User, error) {
	var Users []model.User
	res := d.Db.Find(&Users)
	if res.Error != nil {
		return nil, res.Error
	}
	return Users, nil
}

func (d *DataStore) InsertUser(User *model.User) error {
	return d.Db.Create(User).Error
}

func (d *DataStore) DeleteUserByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) UpdateUser(User *model.User) error {
	err := d.Db.Save(&User).Error
	if err != nil {
		return err
	}
	return nil
}
