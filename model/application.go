package model

import "time"

type Application struct {
	ID          string
	Name        string
	Category    int
	Owner       User `gorm:"foreignKey:OwnerID"`
	Lead        User `gorm:"foreignKey:LeadID"`
	Description string
	JenkinsDir  string `gorm:"jenkins_directory"`
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
