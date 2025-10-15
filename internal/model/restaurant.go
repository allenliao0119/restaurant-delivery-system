package model

import "time"

type Restaurant struct {
	ID        string    `xorm:"pk uuid default gen_random_uuid()"`
	OwnerID   string    `xorm:"uuid notnull index"` // FK -> users.id（以 SQL 方式另加）
	Name      string    `xorm:"varchar(100) notnull index"`
	Address   string    `xorm:"text"`
	Phone     string    `xorm:"varchar(20)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
