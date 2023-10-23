package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestSetClause(t *testing.T) {
	t.Run("should generate set", func(t *testing.T) {
		cases := map[string][]goquent.S{
			"SET status = ?":                    {{"status", "bar"}},
			"SET status = ?, name = ?":          {{"status", "bar"}, {"name", "lucas"}},
			"SET price = ?, cost = ?, name = ?": {{"price", "bar"}, {"cost", "lucas"}, {"name", "productA"}},
		}

		for w, tests := range cases {
			sql, _, err := goquent.New(goquent.MYSQL).Set(tests...).Build()
			if err != nil {
				t.Error(err)
			}

			if w != sql {
				t.Errorf("wanted %s but gotted %s", w, sql)
			}
		}
	})
}
