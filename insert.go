package goquent

type InsertClause struct {
	table string
	q     *QueryBuilder
	f     string
}

func (u *InsertClause) ToSQL() (string, []interface{}, error) {
	return "INSERT INTO " + u.table + "(" + u.f + ")", nil, nil
}

func NewInsertClause(q *QueryBuilder, table, fields string) Clause {
	return &InsertClause{
		q:     q,
		table: table,
		f:     fields,
	}
}
