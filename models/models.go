package models

import "time"

type Person struct {
	ID        uint   `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Gender    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}