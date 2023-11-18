package model

import "time"

type User struct {
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

// DTO input

type AddUserIn struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname"`
	Username  string `json:"username" binding:"required"`
	Profile   int    `json:"profile" binding:"required"`
	Email     string `json:"email"`
}
