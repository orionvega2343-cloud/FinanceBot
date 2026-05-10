package db

import (
	"LangBot/Internal/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.DB) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s user=%s  port=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Port, cfg.Password, cfg.Name, cfg.SslMode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
