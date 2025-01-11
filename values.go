package goquent

import (
	"fmt"
	"strings"
)

type V []interface{}

type ValuesClause struct {
	q      *QueryBuilder
	values []V
}

const valuesClause = "values"

func (ic *ValuesClause) ToSQL() (string, []interface{}, error) {
	sql := "VALUES "
	args := make([]interface{}, 0)

	for i, value := range ic.values {
		if i != 0 {
			sql += ", "
		}

		if ic.q.GetDialect() == MYSQL {
			sql += "(?" + strings.Repeat(", ?", len(value)-1) + ")"
			for _, v := range value {
				args = append(args, v)
			}
		} else if ic.q.GetDialect() == PGSQL {
			sql += "("
			for j, v := range value {
				ic.q.updateIndexPGSQL()

				if j != 0 {
					sql += ", "
				}
				sql += fmt.Sprintf("$%d", ic.q.getIndexPGSQL())

				args = append(args, v)
			}
			sql += ")"
		}
	}

	return sql, args, nil
}

func (c *ValuesClause) WhoIAm() string {
	return valuesClause
}

func NewValuesClause(q *QueryBuilder, values ...V) Clause {
	return &ValuesClause{
		q:      q,
		values: values,
	}
}
