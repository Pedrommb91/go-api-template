package repositories

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
)

type GetUsersMapper struct{}

func NewGetUsersMapper() *GetUsersMapper {
	return &GetUsersMapper{}
}

func (m *GetUsersMapper) Map(rows *sql.Rows) (*openapi.GetUsersResponse, error) {
	const op errors.Op = "repositories.Map"

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
