package model

import "time"

type Application struct {
	ID          string
	Name        string
	Category    int
	Owner       User `gorm:"foreignKey:OwnerID"`
	Lead        User `gorm:"foreignKey:LeadID"`
	Description string
	Created     time.Time
	Updated     time.Time
	OwnerID     int
	LeadID      int
}
