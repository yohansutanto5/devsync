package model

import "time"

type Release struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Application Application `gorm:"foreignKey:AppID"`
	AppID       int
	Status      string
	VersionUAT  string
	VersionPRD  string
	Workflow    string
	Job         string

	Created time.Time
	Updated time.Time
}
