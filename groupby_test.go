package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestGroupByClasure(t *testing.T) {
	t.Run("should generate a group by sql", func(t *testing.T) {
		cases := map[string][]string{
			"GROUP BY name":                 []string{"name"},
			"GROUP BY user_id, contract_id": []string{"user_id", "contract_id"},
			"GROUP BY user_id":              []string{"user_id"},
		}

		for wanted, cols := range cases {
			q := goquent.New().GroupBy(cols...)

			sql, _, err := q.Build()
			if err != nil {
				t.Error(err)
			}

			if wanted != sql {
				t.Errorf("wanted %s gotted %s", wanted, sql)
			}
		}
	})

	t.Run("shoul run with outher clausres", func(t *testing.T) {
		q := goquent.New().Select("user_id", "count(1)").From("contracts").GroupBy("user_id")

		sql, _, err := q.Build()
		if err != nil {
			t.Error(err)
		}

		wanted := "SELECT user_id, count(1) FROM contracts GROUP BY user_id"
		if sql != wanted {
			t.Errorf("wanted %s but gotted %s", wanted, sql)
		}
	})
}
