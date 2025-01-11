package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateSelectForPagination(t *testing.T) {
	t.Run("should generate query and count", func(t *testing.T) {
		sql, countSql, _, _, err := goquent.New(goquent.MYSQL).
			Select("name, email, status, created").
			From("users").
			Where(goquent.C{"name", "LIKE", "lucas"}).
			BuildForPagination()

		if err != nil {
			t.Errorf("getted error %e", err)
		}

		aim := "SELECT name, email, status, created FROM users WHERE name LIKE ?"
		if sql != aim {
			t.Errorf("wanted %s but getted %s", aim, sql)
		}

		aimCount := "SELECT COUNT(1) FROM users WHERE name LIKE ?"
		if countSql != aimCount {
			t.Errorf("wanted %s but getted %s", aimCount, countSql)
		}
	})

	t.Run("should ignore order by in pagination", func(t *testing.T) {
		sql, countSql, _, _, err := goquent.New(goquent.PGSQL).
			Select("name, email, status, created").
			From("users").
			Where(goquent.C{"status", "=", "actived"}).
			OrderBy("id").
			BuildForPagination()

		if err != nil {
			t.Errorf("getted error %e", err)
		}

		aim := "SELECT name, email, status, created FROM users WHERE status = $1 ORDER BY id ASC"
		if sql != aim {
			t.Errorf("wanted %s but getted %s", aim, sql)
		}

		aimCount := "SELECT COUNT(1) FROM users WHERE status = $1"
		if countSql != aimCount {
			t.Errorf("wanted %s but getted %s", aimCount, countSql)
		}
	})

	t.Run("should ignore limit and offset in pagination", func(t *testing.T) {
		sql, countSql, _, _, err := goquent.New(goquent.PGSQL).
			Select("name, email, status, created").
			From("users").
			Where(goquent.C{"status", "=", "actived"}).
			Limit("10").
			Offset("0").
			BuildForPagination()

		if err != nil {
			t.Errorf("getted error %e", err)
		}

		aim := "SELECT name, email, status, created FROM users WHERE status = $1 LIMIT $2 OFFSET $3"
		if sql != aim {
			t.Errorf("wanted %s but getted %s", aim, sql)
		}

		aimCount := "SELECT COUNT(1) FROM users WHERE status = $1"
		if countSql != aimCount {
			t.Errorf("wanted %s but getted %s", aimCount, countSql)
		}
	})

	t.Run("should dont args of limit in countArgs", func(t *testing.T) {
		_, _, args, countArgs, err := goquent.New(goquent.PGSQL).
			Select("name, email, status, created").
			From("users").
			Where(goquent.C{"status", "=", "actived"}).
			Limit("10").
			Offset("0").
			BuildForPagination()

		if err != nil {
			t.Errorf("getted error %e", err)
		}

		if len(args) != 3 {
			t.Errorf("args need 3 but getted %d", len(args))
		}

		if len(countArgs) != 1 {
			t.Errorf("args need 1 but getted %d", len(countArgs))
		}
	})
}
