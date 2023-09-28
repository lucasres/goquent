package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q := goquent.NewQueryBuilder().Select("user", "email")

	fmt.Print(q.Build())
}
