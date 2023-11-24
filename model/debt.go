package model

import "time"

type Debt struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	FirstName string
	LastName  string
	Profile   UserProfile `gorm:"foreignKey:ProfileID"`
	ProfileID int         // Foreign key
	Email     string
	Active    bool
	Created   time.Time
	Updated   time.Time
}
