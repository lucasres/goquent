package goquent

type Instruction interface {
	ToSQL() string
}