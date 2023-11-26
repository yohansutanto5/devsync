package model

type UserProfile struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:VARCHAR(20);not null;"`
}

// DTO input

type AddUserProfileIn struct {
	Name string `json:"Name" binding:"required"`
}
