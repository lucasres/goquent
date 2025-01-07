package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateOffsetQuery(t *testing.T) {
	t.Run("should generate PGSQL OFFSET query", func(t *testing.T) {
		var aimLimit int = 10
		var aimOffset int = 5
		sql, args, err := goquent.New(goquent.PGSQL).
			Select().
			From("users").
			Limit(aimLimit).
			Offset(aimOffset).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users LIMIT $1 OFFSET $2"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[0].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimLimit {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimLimit)
		}

		argVal, ok = args[1].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimOffset {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimLimit)
		}
	})

	t.Run("should generate MYSQL OFFSET query", func(t *testing.T) {
		var aimLimit int = 10
		var aimOffset int = 5
		sql, args, err := goquent.New(goquent.MYSQL).
			Select().
			From("users").
			Limit(aimLimit).
			Offset(aimOffset).
			Build()

		if err != nil {
			t.Error("dont return error in build process")
		}

		aim := "SELECT * FROM users LIMIT ? OFFSET ?"
		if sql != aim {
			t.Errorf("sql generate wrong: %s but should be %s", sql, aim)
		}

		argVal, ok := args[0].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimLimit {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimLimit)
		}

		argVal, ok = args[1].(int)
		if !ok {
			t.Error("wrong type arg")
		}

		if argVal != aimOffset {
			t.Errorf("wrong bind: %d but should be %d", argVal, aimOffset)
		}
	})
}
