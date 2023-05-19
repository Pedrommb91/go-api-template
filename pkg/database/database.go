package database

import (
	"database/sql"
	"fmt"

	"github.com/Pedrommb91/go-api-template/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresOrDie(cfg config.Database) *PostgresDB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.Schema,
		cfg.SslMode)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return &PostgresDB{DB: db}
}
