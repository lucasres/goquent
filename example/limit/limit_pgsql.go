package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q := goquent.New(goquent.PGSQL).
		Select("name", "email").
		From("users").
		Limit("$1", 10)

	sql, args, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Print(sql)  // SELECT name, email FROM users LIMIT $1
	fmt.Print(args) // []interface{10}
}
