package database

import (
	"database/sql"
	"fmt"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"

	_ "github.com/lib/pq"
)

func OpenSql(cfg config.Database) (*sql.DB, error) {
	const op errors.Op = "database.OpenSql"
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
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithError(err),
			errors.WithMessage("Failed to connect to database"),
			errors.WithSeverity(zerolog.FatalLevel),
		)
	}

	return db, nil
}
