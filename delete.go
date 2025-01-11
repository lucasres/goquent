package goquent

type DeleteClause struct {
	q *QueryBuilder
}

const deleteClause = "delete"

func (u *DeleteClause) ToSQL() (string, []interface{}, error) {
	return "DELETE", nil, nil
}

func (c *DeleteClause) WhoIAm() string {
	return deleteClause
}

func NewDeleteClause(q *QueryBuilder) Clause {
	return &DeleteClause{
		q: q,
	}
}
