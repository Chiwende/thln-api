package models

import "time"

type HeroItem struct {
	ID          int       `gorm:"primaryKey"`
	Heading     string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	ImageUrl    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
