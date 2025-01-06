package goquent

import "strings"

type WhereClause struct {
	conditionals []Conditional
	q            *QueryBuilder
}

func (w *WhereClause) ToSQL() (string, []interface{}, error) {
	if len(w.conditionals) == 0 {
		return "", nil, nil
	}

	args := make([]interface{}, 0)
	lastIndex := len(w.conditionals) - 1
	sqls := make([]string, 0)

	for i, c := range w.conditionals {
		conditioalSql, nextConector, conditioalArgs, err := c.Parse(w.q)
		if err != nil {
			return "", nil, err
		}

		if i != lastIndex {
			conditioalSql = conditioalSql + " " + nextConector
		}

		sqls = append(sqls, conditioalSql)

		if len(conditioalArgs) > 0 {
			for _, a := range conditioalArgs {
				args = append(args, a)
			}
		}
	}

	return "WHERE " + strings.Join(sqls, " "), args, nil
}

func NewWhereClause(q *QueryBuilder, conditionals ...Conditional) Clause {
	return &WhereClause{
		conditionals: conditionals,
		q:            q,
	}
}
