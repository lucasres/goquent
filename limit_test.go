package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateLimitSQL(t *testing.T) {
	t.Run("should generate limit sql in PGSQL", func(t *testing.T) {
		var aimArgs int = 10
		sql, args, err := goquent.New(goquent.PGSQL).
			Select().
			From("users").
			Limit(aimArgs).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users LIMIT $1"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[0].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimArgs {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimArgs)
		}
	})

	t.Run("should generate limit sql in MYSQL", func(t *testing.T) {
		var aimArgs int = 10
		sql, args, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			Limit(aimArgs).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users LIMIT ?"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[0].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimArgs {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimArgs)
		}
	})

	t.Run("should generate limit with where", func(t *testing.T) {
		var aimArgs int = 10
		sql, args, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			Where(goquent.C{"status", "=", "active"}).
			Limit(aimArgs).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users WHERE status = ? LIMIT ?"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[1].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimArgs {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimArgs)
		}
	})

	t.Run("should generate limit and offset query", func(t *testing.T) {
		var aimLimit int = 10
		var aimOffset int = 5
		sql, args, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			Where(goquent.C{"status", "=", "active"}).
			Limit(aimLimit, aimOffset).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users WHERE status = ? LIMIT ?, ?"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[1].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimLimit {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimLimit)
		}

		argVal, ok = args[2].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimOffset {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimOffset)
		}
	})
}
