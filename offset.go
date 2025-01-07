package goquent

type OffsetClause struct {
	q *QueryBuilder
	a []interface{}
}

func (o *OffsetClause) ToSQL() (string, []interface{}, error) {
	var limitStr string
	var args = make([]interface{}, 0)
	limitStr = "OFFSET " + getBindIdentifier(o.q)
	args = o.a

	return limitStr, args, nil
}

func NewOffsetClause(q *QueryBuilder, args ...interface{}) Clause {
	return &OffsetClause{
		q: q,
		a: args,
	}
}
