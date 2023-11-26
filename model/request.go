package model

import "time"

type Request struct {
	ID                  string `gorm:"type:VARCHAR(10);not null;primaryKey"`
	Type                string `gorm:"type:VARCHAR(50);not null"`
	Status              string `gorm:"type:VARCHAR(50);not null"`
	PrimaryApprover     User   `gorm:"foreignKey:PrimaryApproverID"`
	SecondaryApprover   User   `gorm:"foreignKey:SecondaryApproverID"`
	PrimaryApproverID   int
	SecondaryApproverID int
	Application         Application `gorm:"foreignKey:AppID"`
	AppID               string      `gorm:"type:VARCHAR(10);not null;"`
	ApprovalURI         string      `gorm:"type:VARCHAR(150)"`
	Created             time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Updated             time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}