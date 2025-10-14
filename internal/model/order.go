package model

import "time"

type Order struct {
	ID           string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID   string  `gorm:"type:uuid"`
	RestaurantID string  `gorm:"type:uuid"`
	DeliveryID   *string `gorm:"type:uuid"`
	Status       string  `gorm:"type:varchar(20)"`
	TotalPrice   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
