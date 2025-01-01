package goquent

import (
	"fmt"
)

type LimitClause struct {
	q *QueryBuilder
	a []interface{}
}

func (l *LimitClause) ToSQL() (string, []interface{}, error) {
	var limitStr string
	args := make([]interface{}, 0)
	if len(l.a) == 0 {
		return "", nil, fmt.Errorf("in mysql need to set limit value")
	}

	if len(l.a) == 1 && l.q.GetDialect() == MYSQL {
		limitStr = "LIMIT ?"
		args = append(args, l.a[0])
	}

	if len(l.a) > 1 {
		limitStr = fmt.Sprintf("LIMIT %s", l.a[0])
		args = l.a[1:]
	}

	return limitStr, args, nil
}

func NewLimitClause(q *QueryBuilder, args ...interface{}) Clause {
	return &LimitClause{
		q: q,
		a: args,
	}
}
