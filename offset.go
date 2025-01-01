package goquent

import (
	"fmt"
)

type OffsetClause struct {
	q *QueryBuilder
	a []interface{}
}

func (o *OffsetClause) ToSQL() (string, []interface{}, error) {
	var limitStr string
	var args = make([]interface{}, 0)
	if o.q.GetDialect() == MYSQL {
		limitStr = "OFFSET ?"
		args = o.a
	} else {
		limitStr = fmt.Sprintf("OFFSET %s", o.a[0])
		args = o.a[1:]
	}

	return limitStr, args, nil
}

func NewOffsetClause(q *QueryBuilder, args ...interface{}) Clause {
	return &OffsetClause{
		q: q,
		a: args,
	}
}
