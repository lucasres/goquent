package goquent

import "fmt"

type FromInstrution struct {
	table string
}

func (f *FromInstrution) ToSQL() string {
	return fmt.Sprintf("FROM %s", f.table)
}

func NewFromInstruction(t string) Instruction {
	return &FromInstrution{
		table: t,
	}
}
