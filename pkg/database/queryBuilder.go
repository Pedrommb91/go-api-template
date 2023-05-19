package database

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
)

type queryBuilder[T any] struct {
	db       *sql.DB
	queryStr string
}

func Where[T any](db *sql.DB) *queryBuilder[T] {
	b := &queryBuilder[T]{}
	b.db = db
	return b
}

func (q *queryBuilder[T]) Select(field string) *queryBuilder[T] {
	q.queryStr += "SELECT " + field
	return q
}

func (q *queryBuilder[T]) From(table string) *queryBuilder[T] {
	q.queryStr += "FROM " + table
	return q
}

func (q *queryBuilder[T]) Run(mapper func(rows *sql.Rows) (T, error)) ([]T, error) {
	const op errors.Op = "database.Run"

	rows, err := q.db.Query(q.queryStr)
	if err != nil {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithMessage("Failed to get users from database"),
			errors.WithError(err),
			errors.WithSeverity(zerolog.ErrorLevel),
		)
	}

	data := make([]T, 0)
	for rows.Next() {
		element, err := mapper(rows)
		if err != nil {
			return nil, errors.Build(
				errors.WithOp(op),
				errors.WithNestedErrorCopy(err),
			)
		}
		data = append(data, element)
	}
	return data, nil
}
