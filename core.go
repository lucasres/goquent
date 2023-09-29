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

func (q *QueryBuilder) Build() (string, []interface{}) {
	instructionSQL := make([]string, 0)

	for _, v := range q.instructions {
		instructionSQL = append(instructionSQL, v.ToSQL())
	}

	return strings.Join(instructionSQL, " "), q.args
}

func (q *QueryBuilder) appendInstruction(i Instruction) {
	q.instructions = append(q.instructions, i)
}

func New() *QueryBuilder {
	return &QueryBuilder{
		instructions: make([]Instruction, 0),
	}
}
