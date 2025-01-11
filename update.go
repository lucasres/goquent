package goquent

type UpdateClause struct {
	table string
	q     *QueryBuilder
}

const updateClause = "update"

func (u *UpdateClause) ToSQL() (string, []interface{}, error) {
	return "UPDATE " + u.table, nil, nil
}

func (c *UpdateClause) WhoIAm() string {
	return updateClause
}

func NewUpdateClause(q *QueryBuilder, table string) Clause {
	return &UpdateClause{
		q:     q,
		table: table,
	}
}
