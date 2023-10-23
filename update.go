package goquent

type UpdateClause struct {
	table string
	q     *QueryBuilder
}

func (u *UpdateClause) ToSQL() (string, []interface{}, error) {
	return "UPDATE " + u.table, nil, nil
}

func NewUpdateClause(q *QueryBuilder, table string) Clause {
	return &UpdateClause{
		q:     q,
		table: table,
	}
}
