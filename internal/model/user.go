package model

import "time"

type User struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	Role         string `gorm:"type:varchar(20)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
