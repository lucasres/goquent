package goquent

import (
	"fmt"
	"strings"
)

type GroupByClause struct {
	cols []string
}

func (g *GroupByClause) ToSQL() (string, []interface{}, error) {
	sql := fmt.Sprintf("GROUP BY %s", strings.Join(g.cols, ", "))

	return sql, nil, nil
}

func NewGroupByClause(c ...string) Clause {
	return &GroupByClause{
		cols: c,
	}
}
