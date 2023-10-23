package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	g := goquent.New(goquent.MYSQL).
		Update("users").
		Set(
			goquent.S{"name", "Lucas"},
			goquent.S{"status", "pedding"},
		).
		Where(goquent.C{"id", "=", 1})

	sql, args, _ := g.Build()
	fmt.Println("SQL " + sql)
	fmt.Println(args)
}
