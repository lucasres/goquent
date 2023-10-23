package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestUpdateClause(t *testing.T) {
	t.Run("should create a sql for update", func(t *testing.T) {
		cases := map[string]string{
			"UPDATE users":    "users",
			"UPDATE contacts": "contacts",
			"UPDATE sales":    "sales",
		}

		for w, table := range cases {
			g := goquent.New(goquent.MYSQL).Update(table)
			sql, _, _ := g.Build()

			if w != sql {
				t.Errorf("wanted %s but gotted %s", w, sql)
			}
		}

	})
}
