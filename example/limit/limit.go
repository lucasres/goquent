package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func PgsqlExample() {
	q := goquent.New(goquent.MYSQL).
		Select("name", "email").
		From("users").
		Limit(10)

	sql, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Print(sql) // SELECT name, email FROM users LIMIT ?
}
