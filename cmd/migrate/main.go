package main

import (
	"log"

	"github.com/allenliao0119/restaurant-delivery-system/internal/model"
	"github.com/allenliao0119/restaurant-delivery-system/pkg/config"
	"github.com/allenliao0119/restaurant-delivery-system/pkg/db"
)

func main() {
	cfg := config.Load()
	engine, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	if err := engine.Sync2(
		new(model.User),
		new(model.Restaurant),
		new(model.Order),
		new(model.OrderItem),
	); err != nil {
		log.Fatalf("xorm sync2 failed: %v", err)
	}

	log.Println("✅ xorm sync2 done")

	_, _ = engine.Exec(`
		ALTER TABLE restaurants
		ADD CONSTRAINT fk_restaurants_owner
		FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE;
	`)
	_, _ = engine.Exec(`
		ALTER TABLE orders
		ADD CONSTRAINT fk_orders_customer
		FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE RESTRICT;
	`)
	_, _ = engine.Exec(`
		ALTER TABLE orders
		ADD CONSTRAINT fk_orders_restaurant
		FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE RESTRICT;
	`)
	_, _ = engine.Exec(`
		ALTER TABLE orders
		ADD CONSTRAINT fk_orders_delivery
		FOREIGN KEY (delivery_id) REFERENCES users(id) ON DELETE SET NULL;
	`)
	_, _ = engine.Exec(`
		ALTER TABLE order_items
		ADD CONSTRAINT fk_order_items_order
		FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE;
	`)

	log.Println("✅ foreign keys ensured (best-effort)")
}
