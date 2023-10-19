package goquent

import "strings"

type QueryBuilder struct {
	Clauses []Clause
	args    []interface{}
	Dialect int
}

const (
	MYSQL = iota
)

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
	q.appendClause(NewSelectClause(q, cols...))

	return q
}

func (q *QueryBuilder) From(t string) *QueryBuilder {
	q.appendClause(NewFromClause(q, t))

	return q
}

func (q *QueryBuilder) Where(conditionals ...conditional) *QueryBuilder {
	i := NewWhereClause(q, conditionals...)

	q.appendClause(i)

	return q
}

func (q *QueryBuilder) GroupBy(c ...string) *QueryBuilder {
	q.appendClause(NewGroupByClause(q, c...))

	return q
}

func (q *QueryBuilder) Build() (string, []interface{}, error) {
	sql := make([]string, 0)

	for _, v := range q.Clauses {
		ClauseSQL, args, err := v.ToSQL()
		if err != nil {
			return "", nil, err
		}

		if len(args) > 0 {
			for _, a := range args {
				q.args = append(q.args, a)
			}
		}

		sql = append(sql, ClauseSQL)
	}

	return strings.Join(sql, " "), q.args, nil
}

func (q *QueryBuilder) GetDialect() int {
	return q.Dialect
}

func (q *QueryBuilder) appendClause(i Clause) {
	q.Clauses = append(q.Clauses, i)
}

func (q *QueryBuilder) appendArgs(args []interface{}) {
	for _, v := range args {
		q.args = append(q.args, v)
	}
}

func New(d int) *QueryBuilder {
	return &QueryBuilder{
		Clauses: make([]Clause, 0),
		args:    make([]interface{}, 0),
		Dialect: d,
	}
}
