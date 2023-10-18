package goquent

import (
	"fmt"
	"strings"
)

type SelectClause struct {
	cols []string
}

func (s *SelectClause) ToSQL() (string, []interface{}, error) {
	colsStr := "*"

	if len(s.cols) != 0 {
		colsStr = strings.Join(s.cols, ", ")
	}

	return fmt.Sprintf("SELECT %s", colsStr), nil, nil
}

func NewSelectClause(cols ...string) Clause {
	return &SelectClause{
		cols: cols,
	}
}
