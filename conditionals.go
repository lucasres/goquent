package goquent

import (
	"fmt"
)

// This respresent conditional in SQL
// ex: quantity > 50 => C{"quantity", ">", "50", "AND"}
type C []interface{}

// This is used if you want group conditioan lin (... AND ...)
type P []interface{}

type conditional interface {
	Parse() (string, string, []interface{}, error)
}

func (c C) Parse() (string, string, []interface{}, error) {
	sql := ""
	size := len(c)
	conector := "AND"
	args := make([]interface{}, 0)

	if size < 3 {
		return sql, conector, args, fmt.Errorf("invalid C args need 3 required but args goted: %d", size)
	}

	if size > 4 {
		return sql, conector, args, fmt.Errorf("invalid C args max 4 but args goted: %d", size)
	}

	if size == 4 {
		conector, ok := c[3].(string)
		if !ok {
			return sql, conector, args, fmt.Errorf("invalid C args the 4 is a conector between conditional and need a string gotted: %x", c[3])
		}
	}

	parsedTest := ""

	column, ok := c[0].(string)
	if !ok {
		return sql, conector, args, fmt.Errorf("in C index 0 is a column and need string")
	}

	op, ok := c[1].(string)
	if !ok {
		return sql, conector, args, fmt.Errorf("in C index 1 is a operator and need string: \">\", \"<\", \"IN\", \"=\", \"LIKE\"...")
	}

	switch op {
	case "BETWEEN":
		testList, ok := c[2].([]string)
		if !ok || len(testList) != 3 {
			return sql, conector, args, fmt.Errorf("operator between in conditional \"%s\" need []string{\"A\", \"AND\", \"B\"}", column)
		}

		parsedTest = fmt.Sprintf("? %s ?", testList[1])
		args = append(args, testList[0])
		args = append(args, testList[2])
	default:
		parsedTest = "?"
		args = append(args, c[2])
	}

	sql = fmt.Sprintf("%s %s %s", column, op, parsedTest)

	return "WHERE " + sql, conector, args, nil
}

func (p *P) Parse() {

}
