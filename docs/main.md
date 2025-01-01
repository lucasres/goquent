# SELECT

For select you can specific columns or all columns

```go
q := goquent.New().
        Select("name", "email").
        From("users")

sql, _, err := q.Build()
if err != nil {
    panic(err)
}
fmt.Print(sql) // SELECT name, email FROM users
```

```go
q := goquent.New().
        Select().
        From("users")

sql, _, err := q.Build()
if err != nil {
    panic(err)
}
fmt.Print(sql) // SELECT * FROM users
```

## WHERE

We have a one clause of WHERE. You can set n `conditional` for make your query

```go
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

fmt.Println(q) // SELECT * FROM users WHERE status = ? AND created_at BETWEEN ? AND ?
fmt.Println(args) // []interface{"2023-01-01 00:00:00", "2023-12-31 23:59:59"}
```

You can see that default operator between conditionals is `AND` but you can change this just seting one more arg in `goquent.C` the index is the operator

```go
q, args, err := goquent.New(goquent.MYSQL).
		Select("*").
		From("users").
		Where(
			&goquent.C{"status", "=", "actived", "OR"},
			&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
		).
		Build()
if err != nil {
    panic(err)
}

fmt.Println(q) // SELECT * FROM users WHERE status = ? OR created_at BETWEEN ? AND ?
fmt.Println(args) // []interface{"2023-01-01 00:00:00", "2023-12-31 23:59:59"}
```

If you want a group conditional you can use this:
```go
q, args, err := goquent.New(goquent.MYSQL).
		Select("*").
		From("users").
		Where(
            &goquent.P{
                &goquent.C{"status", "=", "actived", "OR"},
                &goquent.C{"status", "=", "pedding"},    
            },
			&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
		).
		Build()
if err != nil {
    panic(err)
}

fmt.Println(q) // SELECT * FROM users WHERE (status = ? OR status = ?) AND created_at BETWEEN ? AND ?
fmt.Println(args) // []interface{"2023-01-01 00:00:00", "2023-12-31 23:59:59"}
```

if you want change operator between P and C can you make this:

```go
q, args, err := goquent.New(goquent.MYSQL).
		Select("*").
		From("users").
		Where(
            &goquent.P{
                &goquent.C{"status", "=", "actived", "OR"},
                &goquent.C{"status", "=", "pedding"},
				"OR",
            },
			&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
		).
		Build()
if err != nil {
    panic(err)
}

fmt.Println(q) // SELECT * FROM users WHERE (status = ? OR status = ?) OR created_at BETWEEN ? AND ?
fmt.Println(args) // []interface{"2023-01-01 00:00:00", "2023-12-31 23:59:59"}
```