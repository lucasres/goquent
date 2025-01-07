package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateOrderBytQuery(t *testing.T) {
	t.Run("should generate order by query", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.PGSQL).
			Select().
			From("users").
			OrderBy("name").
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users ORDER BY name ASC"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}
	})

	t.Run("should generate order by query with more fields", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			OrderBy("name", "created_at").
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users ORDER BY name, created_at ASC"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}
	})

	t.Run("should generate order by query with DESC order", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			OrderBy("name", "created_at", "DESC").
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users ORDER BY name, created_at DESC"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}
	})

	t.Run("cant generate order by query", func(t *testing.T) {
		_, _, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			OrderBy().
			Build()

		if err == nil {
			t.Error("must be return error in build process")
		}
	})
}
