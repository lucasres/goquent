package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q := goquent.New().
		Select("name", "email").
		From("users")

	sql, _ := q.Build()
	fmt.Print(sql)
}
