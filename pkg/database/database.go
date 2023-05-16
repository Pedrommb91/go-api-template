package database

import (
	"fmt"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenPostgres(cfg config.Database) (*gorm.DB, error) {
	const op errors.Op = "database.Open"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DbName,
		cfg.Port,
		cfg.Schema,
		cfg.SslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
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
