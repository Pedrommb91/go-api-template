package repositories

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/database"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
)

type PostgresDB struct {
	DB *sql.DB
}

type UsersRepository interface {
	GetUsers() ([]*openapi.GetUsersResponse, error)
}

func (p *PostgresDB) GetUsers() ([]*openapi.GetUsersResponse, error) {
	const op errors.Op = "database.GetUsers"

	users, err := database.Where[*openapi.GetUsersResponse](p.DB).Select("*").From("public.users").WithMapper(NewGetUsersMapper()).Run()
	if err != nil {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithMessage("Failed to get users from database"),
			errors.WithError(err),
			errors.WithSeverity(zerolog.ErrorLevel),
		)
	}

	return users, nil
}
