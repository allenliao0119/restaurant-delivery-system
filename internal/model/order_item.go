package model

type OrderItem struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderID  string `gorm:"type:uuid"`
	Name     string
	Quantity int
	Price    float64
}
