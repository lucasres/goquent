package goquent_test

import (
	"testing"

	"github.com/lucasres/goquent"
)

func TestMain(t *testing.T) {
	t.Run("should dialect set in struct", func(t *testing.T) {
		q := goquent.New(goquent.MYSQL)

		if q.GetDialect() != goquent.MYSQL {
			t.Error("dialect must be set with mysql")
		}
	})
}
