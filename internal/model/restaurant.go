package model

import "time"

type Restaurant struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID   string `gorm:"type:uuid"`
	Name      string
	Address   string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
