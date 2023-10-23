package goquent

import (
	"errors"
	"fmt"
	"strings"
)

type S []interface{}

type SetClause struct {
	sets []S
	q    *QueryBuilder
}

func (s *SetClause) ToSQL() (string, []interface{}, error) {
	sql := make([]string, 0)
	args := make([]interface{}, 0)

	for _, v := range s.sets {
		if len(v) != 2 {
			return "", nil, errors.New("invalid args for S need string, interface{}")
		}

		column, ok := v[0].(string)
		if !ok {
			return "", nil, errors.New("invalid args the index 0 in S need string")
		}

		currentSql := fmt.Sprintf("%s = ?", column)
		args = append(args, v[1])
		sql = append(sql, currentSql)
	}

	return "SET " + strings.Join(sql, ", "), args, nil
}

func NewSetClause(q *QueryBuilder, sets ...S) *SetClause {
	return &SetClause{
		q:    q,
		sets: sets,
	}
}
