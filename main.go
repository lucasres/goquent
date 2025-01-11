package goquent

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type CallbackScanFunc func(rows *sql.Rows) error
type CallbackCountScanFunc func(row *sql.Row) error

type QueryBuilder struct {
	Clauses           []Clause
	args              []interface{}
	Dialect           int
	currentIndexPGSQL int
}

const (
	MYSQL = iota
	PGSQL
)

const (
	PGSQL_CONFLIT_DO_NOTHING = "DO NOTHING"
)

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
	q.appendClause(NewSelectClause(q, cols...))

	return q
}

func (q *QueryBuilder) From(t string) *QueryBuilder {
	q.appendClause(NewFromClause(q, t))

	return q
}

func (q *QueryBuilder) Where(conditionals ...Conditional) *QueryBuilder {
	i := NewWhereClause(q, conditionals...)

	q.appendClause(i)

	return q
}

func (q *QueryBuilder) GroupBy(c ...string) *QueryBuilder {
	q.appendClause(NewGroupByClause(q, c...))

	return q
}

func (q *QueryBuilder) Update(t string) *QueryBuilder {
	q.appendClause(NewUpdateClause(q, t))

	return q
}

func (q *QueryBuilder) Insert(table, fields string) *QueryBuilder {
	q.appendClause(NewInsertClause(q, table, fields))

	return q
}

func (q *QueryBuilder) Values(values ...V) *QueryBuilder {
	q.appendClause(NewValuesClause(q, values...))

	return q
}

func (q *QueryBuilder) Set(sets ...S) *QueryBuilder {
	q.appendClause(NewSetClause(q, sets...))

	return q
}

func (q *QueryBuilder) Limit(args ...interface{}) *QueryBuilder {
	q.appendClause(NewLimitClause(q, args...))

	return q
}

func (q *QueryBuilder) Offset(args ...interface{}) *QueryBuilder {
	q.appendClause(NewOffsetClause(q, args...))

	return q
}

func (q *QueryBuilder) Conflict(values ...string) *QueryBuilder {
	q.appendClause(NewConflictClause(q, values...))

	return q
}

func (q *QueryBuilder) Delete() *QueryBuilder {
	q.appendClause(NewDeleteClause(q))

	return q
}

// Add ORDER BY in query
// separate fields by , EX: OrderBy("name", "created_at")
// if want change order ASC or DESC, set ASC or DESC in last arg ex: OrderBy("name", "created_at", "DESC")
func (q *QueryBuilder) OrderBy(fields ...string) *QueryBuilder {
	q.appendClause(NewOrderByClause(q, fields...))

	return q
}

func (q *QueryBuilder) setClauses(clauses ...Clause) *QueryBuilder {
	q.Clauses = clauses

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

// Returns sql select and count(1) for get total rows
// must used in pagination who want get list and count total rows for pagination
func (q *QueryBuilder) BuildForPagination() (string, string, []interface{}, []interface{}, error) {
	sql := make([]string, 0)
	totalSelect := make([]string, 0)
	totalArgs := make([]interface{}, 0)

	for _, v := range q.Clauses {
		clauseSQL, args, err := v.ToSQL()
		if err != nil {
			return "", "", nil, nil, err
		}

		if len(args) > 0 {
			for _, a := range args {
				q.args = append(q.args, a)
			}
		}

		sql = append(sql, clauseSQL)
		// in pagination dont have order by and select was diferent, need COUNT
		if v.WhoIAm() != orderbyClause && v.WhoIAm() != selectClause && v.WhoIAm() != limitClause && v.WhoIAm() != offsetClause {
			totalSelect = append(totalSelect, clauseSQL)
			if len(args) > 0 {
				for _, a := range args {
					totalArgs = append(totalArgs, a)
				}
			}
		}
		// change selecto to COUNT(1)
		if v.WhoIAm() == selectClause {
			newSelect := NewSelectClause(q, "COUNT(1)")
			newSelectSql, _, _ := newSelect.ToSQL()
			totalSelect = append(totalSelect, newSelectSql)
		}
	}

	return strings.Join(sql, " "), strings.Join(totalSelect, " "), q.args, totalArgs, nil
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

func (q *QueryBuilder) updateIndexPGSQL() {
	q.currentIndexPGSQL += 1
}

func (q *QueryBuilder) getIndexPGSQL() int {
	return q.currentIndexPGSQL
}

func New(d int) *QueryBuilder {
	return &QueryBuilder{
		Clauses:           make([]Clause, 0),
		args:              make([]interface{}, 0),
		Dialect:           d,
		currentIndexPGSQL: 0,
	}
}

func QueryContext(
	ctx context.Context,
	db *sql.DB,
	q *QueryBuilder,
	c CallbackScanFunc,
) error {
	sql, args, err := q.Build()
	if err != nil {
		return fmt.Errorf("build query error: %w", err)
	}

	rows, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec query fail: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		err := c(rows)
		if err != nil {
			return err
		}
	}

	return nil
}

func QueryPaginationContext(
	ctx context.Context,
	db *sql.DB,
	q *QueryBuilder,
	rowsCallback CallbackScanFunc,
	countCallback CallbackCountScanFunc,
) error {
	sql, sqlCount, args, argsCount, err := q.BuildForPagination()
	if err != nil {
		return fmt.Errorf("build query error: %w", err)
	}

	rows, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec query fail: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rowsCallback(rows)
		if err != nil {
			return err
		}
	}

	row := db.QueryRowContext(ctx, sqlCount, argsCount...)
	err = countCallback(row)
	if err != nil {
		return err
	}

	return nil
}
