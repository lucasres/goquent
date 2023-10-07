package goquent

import "strings"

type QueryBuilder struct {
	instructions []Instruction
	args         []interface{}
}

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
	q.appendInstruction(NewSelectInstruction(cols...))

	return q
}

func (q *QueryBuilder) From(t string) *QueryBuilder {
	q.appendInstruction(NewFromInstruction(t))

	return q
}

func (q *QueryBuilder) Where(conditionals ...conditional) *QueryBuilder {
	i := NewWhereInstruction(conditionals...)

	q.appendInstruction(i)

	return q
}

func (q *QueryBuilder) Build() (string, []interface{}, error) {
	sql := make([]string, 0)

	for _, v := range q.instructions {
		instructionSQL, args, err := v.ToSQL()
		if err != nil {
			return "", nil, err
		}

		if len(args) > 0 {
			for _, a := range args {
				q.args = append(q.args, a)
			}
		}

		sql = append(sql, instructionSQL)
	}

	return strings.Join(sql, " "), q.args, nil
}

func (q *QueryBuilder) appendInstruction(i Instruction) {
	q.instructions = append(q.instructions, i)
}

func (q *QueryBuilder) appendArgs(args []interface{}) {
	for _, v := range args {
		q.args = append(q.args, v)
	}
}

func New() *QueryBuilder {
	return &QueryBuilder{
		instructions: make([]Instruction, 0),
		args:         make([]interface{}, 0),
	}
}
