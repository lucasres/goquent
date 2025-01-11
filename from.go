package goquent

import "fmt"

type FromClause struct {
	table string
	q     *QueryBuilder
}

const fromClause = "from"

func (f *FromClause) ToSQL() (string, []interface{}, error) {
	return fmt.Sprintf("FROM %s", f.table), nil, nil
}

func (c *FromClause) WhoIAm() string {
	return conflictClause
}

func NewFromClause(q *QueryBuilder, t string) Clause {
	return &FromClause{
		table: t,
		q:     q,
	}
}
