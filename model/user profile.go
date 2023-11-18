package model

type UserProfile struct {
	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string
}

// DTO input

type AddUserProfileIn struct {
	Name string `json:"Name" binding:"required"`
}
