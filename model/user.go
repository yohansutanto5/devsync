package model

import "time"

type User struct {
	ID        int         `gorm:"primaryKey;autoIncrement"`
	Username  string      `gorm:"type:VARCHAR(20);not null;"`
	FirstName string      `gorm:"type:VARCHAR(70);not null;"`
	LastName  string      `gorm:"type:VARCHAR(70)"`
	Profile   UserProfile `gorm:"foreignKey:ProfileID"`
	ProfileID int         // Foreign key
	Email     string      `gorm:"type:VARCHAR(70);not null;"`
	Active    bool        `gorm:"type:bool;not null;"`
	Created   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
	Updated   time.Time   `gorm:"type:date;default:(CURRENT_DATE)"`
}

// DTO input

type AddUserIn struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname"`
	Username  string `json:"username" binding:"required"`
	Profile   int    `json:"profile" binding:"required"`
	Email     string `json:"email"`
}
