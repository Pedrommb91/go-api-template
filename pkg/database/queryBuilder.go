package database

import (
	"database/sql"
	"fmt"

	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/rs/zerolog"
)

type queryBuilder[T any] struct {
	db       *sql.DB
	queryStr string
	mapper   QueryMapper[T]
}

type QueryMapper[T any] interface {
	Map(rows *sql.Rows) (T, error)
}

func Where[T any](db *sql.DB) *queryBuilder[T] {
	b := &queryBuilder[T]{}
	b.db = db
	return b
}

func (q *queryBuilder[T]) Select(field string) *queryBuilder[T] {
	q.queryStr += " SELECT " + field
	return q
}

func (q *queryBuilder[T]) From(table string) *queryBuilder[T] {
	q.queryStr += " FROM " + table
	return q
}

func (q *queryBuilder[T]) WithMapper(mapper QueryMapper[T]) *queryBuilder[T] {
	q.mapper = mapper
	return q
}

func (q *queryBuilder[T]) Run() ([]T, error) {
	const op errors.Op = "database.Run"
	if !q.queryIsValid() {
		return nil, errors.Build(
			errors.WithOp(op),
			errors.WithError(fmt.Errorf("query is no valid")),
			errors.WithMessage("Query to database invalid"),
			errors.KindBadRequest(),
		)
	}

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
		element, err := q.mapper.Map(rows)
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

func (q *queryBuilder[T]) queryIsValid() bool {
	return q.db != nil && q.queryStr != "" && q.mapper != nil
}
