package goquent

import "fmt"

type FromClause struct {
	table string
}

func (f *FromClause) ToSQL() (string, []interface{}, error) {
	return fmt.Sprintf("FROM %s", f.table), nil, nil
}

func NewFromClause(t string) Clause {
	return &FromClause{
		table: t,
	}
}
