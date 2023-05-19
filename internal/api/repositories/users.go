package repositories

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
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
			return nil, errors.Build(
				errors.WithOp(op),
				errors.WithNestedErrorCopy(err),
			)
		}
		users = append(users, user)
	}

	return users, nil
}

func mapRowsToGetUsersResponse(rows *sql.Rows) (*openapi.GetUsersResponse, error) {
	const op errors.Op = "database.mapRowsToGetUsersResponse"

	user := new(openapi.GetUsersResponse)
	if err := rows.Scan(&user.Id, &user.Name); err != nil {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithError(err),
			errors.WithMessage("Failed to get all users"),
			errors.KindInternalServerError(),
			errors.WithSeverity(zerolog.ErrorLevel),
		)
	}

	return user, nil
}
