package goquent

import (
	"fmt"
	"strings"
)

type SelectClause struct {
	cols []string
	q    *QueryBuilder
}

const selectClause = "select"

func (s *SelectClause) ToSQL() (string, []interface{}, error) {
	colsStr := "*"

	if len(s.cols) != 0 {
		colsStr = strings.Join(s.cols, ", ")
	}

	return fmt.Sprintf("SELECT %s", colsStr), nil, nil
}

func (c *SelectClause) WhoIAm() string {
	return selectClause
}

func NewSelectClause(q *QueryBuilder, cols ...string) Clause {
	return &SelectClause{
		cols: cols,
		q:    q,
	}
}
