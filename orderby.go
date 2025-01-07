package goquent

import (
	"errors"
	"strings"
)

type OrderByClause struct {
	q      *QueryBuilder
	fields []string
}

func (o *OrderByClause) ToSQL() (string, []interface{}, error) {
	var limitStr string
	order := "ASC"

	totalArgs := len(o.fields)

	if totalArgs == 0 {
		return "", nil, errors.New("need set almost one field")
	}

	lastIndex := o.fields[totalArgs-1]

	if lastIndex == "ASC" || lastIndex == "DESC" {
		order = lastIndex
		o.fields = o.fields[:totalArgs-1]
	}

	limitStr = "ORDER BY " + strings.Join(o.fields, ", ") + " " + order

	return limitStr, nil, nil
}

func NewOrderByClause(q *QueryBuilder, fields ...string) Clause {
	return &OrderByClause{
		q:      q,
		fields: fields,
	}
}
