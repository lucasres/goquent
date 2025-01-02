package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateInsert(t *testing.T) {
	t.Run("should generate insert query in MYSQL", func(t *testing.T) {
		value := goquent.V{"Lucas", 1, "email@email.com"}
		sql, args, err := goquent.New(goquent.MYSQL).
			Insert("users", "name, status, email").
			Values(value).
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES (?, ?, ?)"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		validArgs(args, value, t)
	})

	t.Run("should generate multiples insert query in MYSQL", func(t *testing.T) {
		values := []goquent.V{
			{"Foo", 1, "email@email.com"},
			{"Bar", 0, "email2@email.com"},
			{"FooBar", 0, "email3@email.com"},
		}

		sql, args, err := goquent.New(goquent.MYSQL).
			Insert("users", "name, status, email").
			Values(values...).
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?)"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		for i := 0; i < len(args); i += 3 {
			validArgs(args[i:i+3], values[i/3], t)
		}
	})

	t.Run("should generate insert query in PGSQL", func(t *testing.T) {
		value := goquent.V{"Lucas", 1, "email@email.com"}
		sql, args, err := goquent.New(goquent.PGSQL).
			Insert("users", "name, status, email").
			Values(value).
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES ($1, $2, $3)"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		validArgs(args, value, t)
	})

	t.Run("should generate multiples insert query in PGSQL", func(t *testing.T) {
		values := []goquent.V{
			{"Foo", 1, "email@email.com"},
			{"Bar", 0, "email2@email.com"},
			{"FooBar", 0, "email3@email.com"},
		}

		sql, args, err := goquent.New(goquent.PGSQL).
			Insert("users", "name, status, email").
			Values(values...).
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9)"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		for i := 0; i < len(args); i += 3 {
			validArgs(args[i:i+3], values[i/3], t)
		}
	})
}

func validArgs(args []interface{}, value goquent.V, t *testing.T) {
	for i, v := range args {
		if v != value[i] {
			t.Errorf("should return %v but get %v", value[i], v)
		}
	}
}
