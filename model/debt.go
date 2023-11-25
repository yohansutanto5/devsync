package model

import "time"

type Debt struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Application Application `gorm:"foreignKey:AppID"`
	AppID       string      `gorm:"type:VARCHAR(10);not null;"`
	Category    string      `gorm:"not null"`
	Status      string      `gorm:"not null;default:'NEW'"`
	Due         time.Time   `gorm:"type:date;not null"`
	Created     time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Updated     time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}
type InsertDebtIn struct {
	AppID    string    `convert:"AppID"`
	Category string    `convert:"Category"`
	Due      time.Time `convert:"Due"`
}

// Table Repeat Deb is to be consumed by scheduler to add debt every interval for each appcode
type RepeatDebt struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Application Application `gorm:"foreignKey:AppID"`
	AppID       string      `gorm:"not null"`
	Category    string      `gorm:"not null"`
	Status      string      `gorm:"not null;default:'NEW'"`
	Due         time.Time   `gorm:"type:date;not null"`
	Created     time.Time   `gorm:"type:date;not null;default:CURRENT_DATE"`
	Updated     time.Time   `gorm:"type:date;not null;default:CURRENT_DATE"`
}
