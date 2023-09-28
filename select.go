package goquent

import (
	"fmt"
	"strings"
)

type SelectInstruction struct {
	cols []string
}

func (s *SelectInstruction) ToSQL() string {
	colsStr := "*"

	if len(s.cols) != 0 {
		colsStr = strings.Join(s.cols, ", ")
	}

	return fmt.Sprintf("SELECT %s", colsStr)
}

func NewSelectInstruction(cols ...string) Instruction {
	return &SelectInstruction{
		cols: cols,
	}
}
