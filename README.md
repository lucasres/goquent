# Description

Hello this is a lib for query build that sintaxe like a SQL. That aim of lib is create a sintaxe very similar to sql if you know sql you should know how this lib work. 

# Get Start

Get the lib:

```
go get github.com/lucasres/goquent
```

Now you can create a query

```
package main

import (
	"fmt"

	"github.com/lucasres/goquent"
)

func main() {
	q := goquent.New().
		Select("name", "email").
		From("users")

	sql, _, err := q.Build()
	if err != nil {
		panic(err)
	}
	fmt.Print(sql) // SELECT name, email FROM users
}
```


