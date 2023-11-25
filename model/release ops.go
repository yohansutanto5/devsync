package model

import "time"

type ReleaseTicket struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Application Application `gorm:"foreignKey:AppID"`
	AppID       string      `gorm:"type:VARCHAR(10);not null;"`
	Status      string
	VersionUAT  string
	VersionPRD  string
	Workflow    string
	Job         string

	Created time.Time
	Updated time.Time
}

type InsertReleaseTicketIn struct {
	AppID      string
	VersionUAT string
	VersionPRD string
	Workflow   string
	Job        string
}
