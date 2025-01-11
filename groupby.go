package goquent

import (
	"fmt"
	"strings"
)

type GroupByClause struct {
	cols []string
	q    *QueryBuilder
}

const groupbyClause = "groupby"

func (g *GroupByClause) ToSQL() (string, []interface{}, error) {
	sql := fmt.Sprintf("GROUP BY %s", strings.Join(g.cols, ", "))

	return sql, nil, nil
}

func (c *GroupByClause) WhoIAm() string {
	return groupbyClause
}

func NewGroupByClause(q *QueryBuilder, c ...string) Clause {
	return &GroupByClause{
		cols: c,
		q:    q,
	}
}
