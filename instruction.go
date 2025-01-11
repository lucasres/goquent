package goquent

type Clause interface {
	ToSQL() (string, []interface{}, error)
	WhoIAm() string
}
