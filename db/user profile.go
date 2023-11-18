package db

import (
	"app/model"
)

func (d *DataStore) GetListUserProfile() ([]model.UserProfile, error) {
	var UserProfiles []model.UserProfile
	res := d.Db.Find(&UserProfiles)
	if res.Error != nil {
		return nil, res.Error
	}
	return UserProfiles, nil
}

func (d *DataStore) InsertUserProfile(UserProfile *model.UserProfile) error {
	return d.Db.Create(UserProfile).Error
}

func (d *DataStore) DeleteUserProfileByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.UserProfile{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) UpdateUserProfile(UserProfile *model.UserProfile) error {
	err := d.Db.Save(&UserProfile).Error
	if err != nil {
		return err
	}
	return nil
}
