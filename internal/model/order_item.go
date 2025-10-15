package model

type OrderItem struct {
	ID       string  `xorm:"pk uuid default gen_random_uuid()"`
	OrderID  string  `xorm:"uuid notnull index"` // FK -> orders.id
	Name     string  `xorm:"varchar(100) notnull"`
	Quantity int     `xorm:"int notnull"`
	Price    float64 `xorm:"numeric(10,2) notnull default 0"`
}
