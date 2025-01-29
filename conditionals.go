package goquent

import (
	"errors"
	"fmt"
	"strings"
)

// This respresent conditional in SQL
// ex: quantity > 50 => C{"quantity", ">", "50", "AND"}
type C []interface{}

// This is used if you want group your conditionals (... AND ...)
// ex: you should pass n &goquent.C{}
type P []interface{}

type Conditional interface {
	Parse(q *QueryBuilder) (string, string, []interface{}, error)
}

func (c C) Parse(q *QueryBuilder) (string, string, []interface{}, error) {
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
		conectorParsed, ok := c[3].(string)
		if !ok {
			return sql, conector, args, fmt.Errorf("invalid C args the 4 is a conector between conditional and need a string gotted: %x", c[3])
		}

		conector = conectorParsed
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

		parsedTest = fmt.Sprintf("%s %s %s", getBindIdentifier(q), testList[1], getBindIdentifier(q))
		args = append(args, testList[0])
		args = append(args, testList[2])
	case "IN":
		testList, ok := c[2].([]interface{})
		if !ok && len(testList) > 0 {
			return sql, conector, args, fmt.Errorf("operator IN in conditional \"%s\" need []interface{\"A\", \"B\", \"C\", ...}", column)
		}
		lastIndex := len(testList) - 1

		parsedTest += "("

		for i, v := range testList {
			parsedTest += getBindIdentifier(q)
			if i != lastIndex {
				parsedTest += ","
			}
			args = append(args, v)
		}

		parsedTest += ")"

		args = append(args, testList[0])
		args = append(args, testList[2])
	default:
		parsedTest = getBindIdentifier(q)
		args = append(args, c[2])
	}

	sql = fmt.Sprintf("%s %s %s", column, op, parsedTest)

	return sql, conector, args, nil
}

func (p P) Parse(q *QueryBuilder) (string, string, []interface{}, error) {
	sql := make([]string, 0)
	args := make([]interface{}, 0)
	lastIndex := len(p) - 1
	conector := "AND"

	for i, v := range p {
		parsedC, ok := v.(*C)
		if ok {
			sqlC, conectorC, argsC, err := parsedC.Parse(q)
			if err != nil {
				return "", "", nil, err
			}

			if i != lastIndex && p.moreCondintional(i) {
				sqlC = sqlC + " " + conectorC
			}

			for _, a := range argsC {
				args = append(args, a)
			}

			sql = append(sql, sqlC)
			continue
		}

		parsedS, ok := v.(string)
		if ok {
			if i == lastIndex {
				conector = parsedS
				continue
			}

			return "", "", nil, errors.New("only last index is the conector must other should be *goquent.C")
		}

		return "", "", nil, fmt.Errorf("invalid type at index %d should be *goquent.C or string", i)
	}
	return "(" + strings.Join(sql, " ") + ")", conector, args, nil
}

func (p P) moreCondintional(i int) bool {
	nI := i + 1

	if nI > len(p)-1 {
		return false
	}

	_, ok := p[i+1].(*C)
	return ok
}

func getBindIdentifier(q *QueryBuilder) string {
	if q.GetDialect() == MYSQL {
		return "?"
	}

	q.updateIndexPGSQL()
	return fmt.Sprintf("$%d", q.getIndexPGSQL())
}
