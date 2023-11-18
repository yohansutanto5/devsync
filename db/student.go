package db

import (
	"app/model"
)

func (d *DataStore) GetListStudent() ([]model.Student, error) {
	var students []model.Student
	res := d.Db.Find(&students)
	if res.Error != nil {
		return nil, res.Error
	}
	return students, nil
}

func (d *DataStore) InsertStudent(student *model.Student) error {
	return d.Db.Create(student).Error
}

func (d *DataStore) DeleteStudentByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.Student{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) UpdateStudent(student *model.Student) error {
	err := d.Db.Save(&student).Error
	if err != nil {
		return err
	}
	return nil
}
