package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestDeleteClause(t *testing.T) {
	t.Run("should generate delete query in MYSQL", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.MYSQL).
			Delete().
			From("users").
			Where(goquent.C{"id", "=", 1}).
			Build()

		if err != nil {
			t.Errorf("cant return err: %e", err)
		}

		aim := "DELETE FROM users WHERE id = ?"
		if aim != sql {
			t.Errorf("get %s but need %s", sql, aim)
		}
	})

	t.Run("should generate delete query in PGSQL", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.PGSQL).
			Delete().
			From("users").
			Where(goquent.C{"id", "=", 1}).
			Build()

		if err != nil {
			t.Errorf("cant return err: %e", err)
		}

		aim := "DELETE FROM users WHERE id = $1"
		if aim != sql {
			t.Errorf("get %s but need %s", sql, aim)
		}
	})

	t.Run("should generate delete query in PGSQL with more args", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.PGSQL).
			Delete().
			From("users").
			Where(goquent.C{"id", "=", 1}, goquent.C{"status", "=", "alive"}).
			Build()

		if err != nil {
			t.Errorf("cant return err: %e", err)
		}

		aim := "DELETE FROM users WHERE id = $1 AND status = $2"
		if aim != sql {
			t.Errorf("get %s but need %s", sql, aim)
		}
	})
}
