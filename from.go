package goquent

import "fmt"

type FromInstrution struct {
	table string
}

func (f *FromInstrution) ToSQL() (string, []interface{}, error) {
	return fmt.Sprintf("FROM %s", f.table), nil, nil
}

func NewFromInstruction(t string) Instruction {
	return &FromInstrution{
		table: t,
	}
}
