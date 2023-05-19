package database

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
)

//go:generate mockery --name UsersRepository
type UsersRepository interface {
	GetUsers() ([]*openapi.GetUsersResponse, error)
}

func (p *PostgresDB) GetUsers() ([]*openapi.GetUsersResponse, error) {
	const op errors.Op = "database.GetUsers"
	rows, err := p.DB.Query("SELECT * FROM public.users")
	if err != nil {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithMessage("Failed to get users from database"),
			errors.WithError(err),
			errors.WithSeverity(zerolog.ErrorLevel),
		)
	}

	users := []*openapi.GetUsersResponse{}
	for rows.Next() {
		user, err := mapRowsToGetUsersResponse(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func mapRowsToGetUsersResponse(rows *sql.Rows) (*openapi.GetUsersResponse, error) {
	user := new(openapi.GetUsersResponse)
	err := rows.Scan(
		&user.Id,
		&user.Name,
	)
	return user, err
}
