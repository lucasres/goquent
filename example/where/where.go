package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q, args, err := goquent.New(goquent.MYSQL).
		Select("*").
		From("users").
		Where(
			&goquent.C{"status", "=", "actived"},
			&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
		).
		Build()
	if err != nil {
		panic(err)
	}

	fmt.Println(q)
	fmt.Println(args)
}
