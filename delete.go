package goquent

type DeleteClause struct {
	q *QueryBuilder
}

func (u *DeleteClause) ToSQL() (string, []interface{}, error) {
	return "DELETE", nil, nil
}

func NewDeleteClause(q *QueryBuilder) Clause {
	return &DeleteClause{
		q: q,
	}
}
