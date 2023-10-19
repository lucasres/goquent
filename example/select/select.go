package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q := goquent.New(goquent.MYSQL).
		Select("name", "email").
		From("users")

	sql, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Print(sql)
}
