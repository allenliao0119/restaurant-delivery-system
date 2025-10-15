package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/allenliao0119/restaurant-delivery-system/pkg/config"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func Connect(cfg config.Config) (*xorm.Engine, error) {
	// lib/pq DSN
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode,
	)

	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// 連線池參數（可依負載調整）
	engine.SetMaxOpenConns(25)
	engine.SetMaxIdleConns(25)
	engine.SetConnMaxLifetime(30 * time.Minute)

	// 命名轉換：GonicMapper 會把欄位轉成 snake_case（CreatedAt -> created_at）
	engine.SetMapper(names.GonicMapper{})

	// 開發期間建議開啟 SQL 輸出
	engine.ShowSQL(true)

	// ping
	if err := engine.Ping(); err != nil {
		return nil, err
	}

	log.Println("PostgreSQL connected via XORM.")
	return engine, nil
}
