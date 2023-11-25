package model

import "time"

type Application struct {
	ID          string `gorm:"type:VARCHAR(10);not null;primaryKey"`
	Name        string `gorm:"type:VARCHAR(50);not null"`
	Category    int    `gorm:"type:SMALLINT;not null"`
	Owner       User   `gorm:"foreignKey:OwnerID"`
	Lead        User   `gorm:"foreignKey:LeadID"`
	Description string
	JenkinsDir  string `gorm:"jenkins_directory;type:VARCHAR(200)"` // Full Jenkins URL link
	Created     time.Time
	Updated     time.Time
	OwnerID     int
	LeadID      int
}

type InsertApplicationIn struct {
	ID          string `convert:"ID"`
	Name        string `convert:"Name"`
	Category    int    `convert:"Category"`
	Description string `convert:"Description"`
	OwnerID     int    `convert:"OwnerID"`
	LeadID      int    `convert:"LeadID"`
}
