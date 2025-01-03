package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestShouldGenerateSelectForPagination(t *testing.T) {
	t.Run("should generate query and count", func(t *testing.T) {
		sql, countSql, _, err := goquent.New(goquent.MYSQL).
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
}
