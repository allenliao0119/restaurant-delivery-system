package model

import "time"

type User struct {
	ID           string    `xorm:"pk uuid default gen_random_uuid()"`
	Name         string    `xorm:"varchar(100) notnull"`
	Email        string    `xorm:"varchar(100) unique notnull"`
	PasswordHash string    `xorm:"text notnull"`
	Role         string    `xorm:"varchar(20) notnull index"` // customer/restaurant/delivery/admin
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}
