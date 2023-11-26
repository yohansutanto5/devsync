package model

import "time"

type ReleaseTicket struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Application Application `gorm:"foreignKey:AppID"`
	AppID       string      `gorm:"type:VARCHAR(10);not null;"`
	Status      string      `gorm:"type:VARCHAR(30);not null;"`
	VersionUAT  string      `gorm:"type:VARCHAR(10);not null;"`
	VersionPRD  string      `gorm:"type:VARCHAR(10);"`
	Workflow    string      `gorm:"type:VARCHAR(10);not null;"`
	Job         string      `gorm:"type:VARCHAR(100);not null;"`
	
	Created time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Updated time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type InsertReleaseTicketIn struct {
	AppID      string
	VersionUAT string
	VersionPRD string
	Workflow   string
	Job        string
}
