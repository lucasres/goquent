package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateConflictSql(t *testing.T) {
	t.Run("should generate insert query with conflict in PGSQL", func(t *testing.T) {
		value := goquent.V{"Lucas", 1, "email@email.com"}
		sql, args, err := goquent.New(goquent.PGSQL).
			Insert("users", "name, status, email").
			Values(value).
			Conflict(goquent.PGSQL_CONFLIT_DO_NOTHING).
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		validArgs(args, value, t)
	})

	t.Run("should generate insert query with conflict and fields in PGSQL", func(t *testing.T) {
		value := goquent.V{"Lucas", 1, "email@email.com"}
		sql, args, err := goquent.New(goquent.PGSQL).
			Insert("users", "name, status, email").
			Values(value).
			Conflict(goquent.PGSQL_CONFLIT_DO_NOTHING, "email").
			Build()

		if err != nil {
			t.Errorf("cant return error: %e", err)
		}

		aim := "INSERT INTO users(name, status, email) VALUES ($1, $2, $3) ON CONFLICT (email) DO NOTHING"
		if sql != aim {
			t.Errorf("should return %s but get %s", aim, sql)
		}

		validArgs(args, value, t)
	})
}
