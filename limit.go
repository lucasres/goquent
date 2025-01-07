package goquent

import (
	"errors"
	"fmt"
)

type LimitClause struct {
	q *QueryBuilder
	a []interface{}
}

func (l *LimitClause) ToSQL() (string, []interface{}, error) {
	var limitStr string
	args := make([]interface{}, 0)

	if len(l.a) == 1 {
		limitStr = "LIMIT " + getBindIdentifier(l.q)
		args = append(args, l.a[0])
	} else if len(l.a) == 2 {
		limitStr = fmt.Sprintf("LIMIT %s, %s", getBindIdentifier(l.q), getBindIdentifier(l.q))
		args = append(args, l.a[0])
		args = append(args, l.a[1])
	} else {
		return "", nil, errors.New("invalid quantity of args: must be 1 or 2")
	}

	return limitStr, args, nil
}

func NewLimitClause(q *QueryBuilder, args ...interface{}) Clause {
	return &LimitClause{
		q: q,
		a: args,
	}
}
