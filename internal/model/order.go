package model

import "time"

type Order struct {
	ID           string    `xorm:"pk uuid default gen_random_uuid()"`
	CustomerID   string    `xorm:"uuid notnull index"`     // FK -> users.id
	RestaurantID string    `xorm:"uuid notnull index"`     // FK -> restaurants.id
	DeliveryID   *string   `xorm:"uuid null index"`        // FK -> users.id (delivery)
	Status       string    `xorm:"varchar(20) notnull index"` // pending/confirmed/delivering/completed/cancelled
	TotalPrice   float64   `xorm:"numeric(10,2) notnull default 0"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}
