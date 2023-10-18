package goquent

import "strings"

type WhereClause struct {
	conditionals []conditional
}

func (w *WhereClause) ToSQL() (string, []interface{}, error) {
	args := make([]interface{}, 0)
	lastIndex := len(w.conditionals) - 1
	sqls := make([]string, 0)

	for i, c := range w.conditionals {
		conditioalSql, nextConector, conditioalArgs, err := c.Parse()
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

func NewWhereClause(conditionals ...conditional) Clause {
	return &WhereClause{
		conditionals: conditionals,
	}
}
