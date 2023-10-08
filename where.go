package goquent

import "strings"

type WhereInstruction struct {
	conditionals []conditional
}

func (w *WhereInstruction) ToSQL() (string, []interface{}, error) {
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

func NewWhereInstruction(conditionals ...conditional) Instruction {
	return &WhereInstruction{
		conditionals: conditionals,
	}
}
