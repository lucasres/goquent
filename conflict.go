package goquent

import (
	"errors"
	"strings"
)

type ConflictClause struct {
	values []string
	q      *QueryBuilder
	f      string
}

const conflictClause = "conflict"

func (c *ConflictClause) ToSQL() (string, []interface{}, error) {
	if c.q.GetDialect() == PGSQL {
		sql := "ON CONFLICT"
		if len(c.values) > 1 {
			sql += " (" + strings.Join(c.values[1:], ", ") + ")"
		}

		if c.values[0] != PGSQL_CONFLIT_DO_NOTHING {
			return "", nil, errors.New("first args must be action of conflict: example DO NOTHING")
		}

		sql += " " + c.values[0]

		return sql, nil, nil
	}

	return "", nil, errors.New("conflict in MYSQL is not implemented")
}

func (c *ConflictClause) WhoIAm() string {
	return conflictClause
}

func NewConflictClause(q *QueryBuilder, values ...string) Clause {
	return &ConflictClause{
		q:      q,
		values: values,
	}
}
