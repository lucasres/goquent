package goquent

type InsertClause struct {
	table string
	q     *QueryBuilder
	f     string
}

const insertClause = "insert"

func (u *InsertClause) ToSQL() (string, []interface{}, error) {
	return "INSERT INTO " + u.table + "(" + u.f + ")", nil, nil
}

func (c *InsertClause) WhoIAm() string {
	return insertClause
}

func NewInsertClause(q *QueryBuilder, table, fields string) Clause {
	return &InsertClause{
		q:     q,
		table: table,
		f:     fields,
	}
}
