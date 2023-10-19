package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestFromClause(t *testing.T) {
	t.Run("should generate FROM sql", func(t *testing.T) {
		cases := map[string]string{
			"users":    "FROM users",
			"emails":   "FROM emails",
			"products": "FROM products",
		}

		for table, expected := range cases {
			sql, args, err := goquent.New(goquent.MYSQL).From(table).Build()
			if err != nil {
				t.Error(err)
			}

			if len(args) > 0 {
				t.Error("size of args should be 0")
			}

			if sql != expected {
				t.Error("size of args should be 0")
			}
		}
	})
}
