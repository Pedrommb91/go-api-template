package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/ent"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"

	_ "github.com/lib/pq" // postgres driver - import needed
)

// OpenPostgres
// Must close client
func OpenPostgres(cfg config.Database) (*ent.Client, error) {
	const op errors.Op = "database.Open"
	conStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.Password)
	client, err := ent.Open(dialect.Postgres, conStr)
	if err != nil {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithError(err),
			errors.WithMessage("Failed to open database"),
			errors.WithSeverity(zerolog.FatalLevel),
		)
	}
	return client, nil
}
