package db

import (
	"app/model"
)

func (d *DataStore) GetListReleaseTicket() ([]model.ReleaseTicket, error) {
	var ReleaseTickets []model.ReleaseTicket
	res := d.Db.Find(&ReleaseTickets)
	if res.Error != nil {
		return nil, res.Error
	}
	return ReleaseTickets, nil
}

func (d *DataStore) InsertReleaseTicket(ReleaseTicket *model.ReleaseTicket) error {
	return d.Db.Create(ReleaseTicket).Error
}

func (d *DataStore) DeleteReleaseTicketByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.ReleaseTicket{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *DataStore) GetReleaseTicketByID(id int) (ticket *model.ReleaseTicket, err error) {
	res := d.Db.Preload("Application").First(&ticket, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return
}

func (d *DataStore) UpdateReleaseTicket(ReleaseTicket *model.ReleaseTicket) error {
	err := d.Db.Save(&ReleaseTicket).Error
	if err != nil {
		return err
	}
	return nil
}
