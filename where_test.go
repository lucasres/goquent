package goquent_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/lucasres/goquent"
)

func TestWhereClause(t *testing.T) {
	t.Run("should create a where clause", func(t *testing.T) {
		sql, args, err := goquent.New(goquent.MYSQL).Where(goquent.C{"status", "=", "actived"}).Build()
		if err != nil {
			t.Error(err)
		}

		if len(args) != 1 {
			t.Error(errors.New("args must return 1 value"))
		}

		testCase := "WHERE status = ?"
		if sql != testCase {
			t.Error(fmt.Errorf("the sql must be \"%s\" gotted \"%s\"", testCase, sql))
		}
	})

	t.Run("should create with more conditionals", func(t *testing.T) {
		sql, args, err := goquent.New(goquent.MYSQL).Where(
			goquent.C{"status", "=", "actived"},
			goquent.C{"created_at", ">", "2023-12-12"},
		).Build()
		if err != nil {
			t.Error(err)
		}

		if len(args) != 2 {
			t.Error(fmt.Errorf("args must return 2 value gotted %d", len(args)))
		}

		testCase := "WHERE status = ? AND created_at > ?"
		if sql != testCase {
			t.Error(fmt.Errorf("the sql must be \"%s\" gotted \"%s\"", testCase, sql))
		}
	})

	t.Run("should can create with OR operator", func(t *testing.T) {
		sql, args, err := goquent.New(goquent.MYSQL).Where(
			goquent.C{"status", "=", "actived", "OR"},
			goquent.C{"created_at", ">", "2023-12-12"},
			goquent.C{"deleted_at", "IS", "NULL"},
		).Build()
		if err != nil {
			t.Error(err)
		}

		if len(args) != 3 {
			t.Error(fmt.Errorf("args must return 3 value gotted %d", len(args)))
		}

		testCase := "WHERE status = ? OR created_at > ? AND deleted_at IS ?"
		if sql != testCase {
			t.Error(fmt.Errorf("the sql must be \"%s\" gotted \"%s\"", testCase, sql))
		}
	})

	t.Run("should return correct args", func(t *testing.T) {
		argsTestCase := []string{"actived", "2023-12-12", "NULL"}

		_, args, err := goquent.New(goquent.MYSQL).Where(
			goquent.C{"status", "=", argsTestCase[0]},
			goquent.C{"created_at", ">", argsTestCase[1]},
			goquent.C{"deleted_at", "IS", argsTestCase[2]},
		).Build()
		if err != nil {
			t.Error(err)
		}

		if len(args) != 3 {
			t.Error(fmt.Errorf("args must return 3 value gotted %d", len(args)))
		}

		for i, v := range args {
			if v != argsTestCase[i] {
				t.Errorf("value of args at index must be \"%s\" and gotted \"%s\"", argsTestCase[i], v)
			}
		}
	})

	t.Run("should return error with invalid args", func(t *testing.T) {
		_, _, err := goquent.New(goquent.MYSQL).Where(
			goquent.C{"status", "=", "actived", "OR", "bar", "foo"},
		).Build()
		if err == nil {
			t.Error(errors.New("invalid args must return an error"))
		}

		_, _, err = goquent.New(goquent.MYSQL).Where(
			goquent.C{"status", "="},
		).Build()
		if err == nil {
			t.Error(errors.New("invalid args must return an error"))
		}
	})
}

func TestBetweenOperator(t *testing.T) {
	t.Run("should can handle between sql", func(t *testing.T) {
		testArgsCases := []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}

		sql, args, err := goquent.New(goquent.MYSQL).Where(
			goquent.C{"created_at", "BETWEEN", testArgsCases},
		).Build()
		if err != nil {
			t.Error(err)
		}

		if len(args) != 2 {
			t.Error(fmt.Errorf("args must return 2 value gotted %d", len(args)))
		}

		testCase := "WHERE created_at BETWEEN ? AND ?"
		if sql != testCase {
			t.Error(fmt.Errorf("the sql must be \"%s\" gotted \"%s\"", testCase, sql))
		}

		if testArgsCases[0] != args[0] {
			t.Error(fmt.Errorf("the args value must be \"%s\" gotted \"%s\"", testArgsCases[0], args[0]))
		}

		if testArgsCases[2] != args[1] {
			t.Error(fmt.Errorf("the args value must be \"%s\" gotted \"%s\"", testArgsCases[0], args[0]))
		}
	})
}

func TestParameter(t *testing.T) {
	t.Run("should can handle with parameter sql", func(t *testing.T) {
		cases := map[string][]goquent.Conditional{
			"SELECT * FROM users WHERE (status = ? OR status = ?) AND created_at BETWEEN ? AND ?": {
				&goquent.P{
					&goquent.C{"status", "=", "actived", "OR"},
					&goquent.C{"status", "=", "pedding"},
				},
				&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
			},
			"SELECT * FROM users WHERE created_at BETWEEN ? AND ? AND (name = ? OR bar = ?) AND email LIKE ?": {
				&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
				&goquent.P{
					&goquent.C{"name", "=", "Lucas", "OR"},
					&goquent.C{"bar", "=", "foo"},
				},
				&goquent.C{"email", "LIKE", `%email.com%`},
			},
		}

		for wanted, c := range cases {
			sql, _, err := goquent.New(goquent.MYSQL).
				Select("*").
				From("users").
				Where(c...).
				Build()
			if err != nil {
				t.Error(err)
			}

			if sql != wanted {
				t.Errorf("want sql %s but gotted %s", wanted, sql)
			}
		}
	})

	t.Run("should can handle with parameter sql and custom operator", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.MYSQL).
			Select("*").
			From("users").
			Where(
				&goquent.P{
					&goquent.C{"status", "=", "actived", "OR"},
					&goquent.C{"status", "=", "pedding"},
					"OR",
				},
				&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
			).
			Build()
		if err != nil {
			t.Error(err)
		}

		testSql := "SELECT * FROM users WHERE (status = ? OR status = ?) OR created_at BETWEEN ? AND ?"
		if sql != testSql {
			t.Errorf("want sql %s but gotted %s", testSql, sql)
		}
	})

	t.Run("should have a err when index of operator will wrong", func(t *testing.T) {
		_, _, err := goquent.New(goquent.MYSQL).
			Select("*").
			From("users").
			Where(
				&goquent.P{
					&goquent.C{"status", "=", "actived", "OR"},
					"OR",
					&goquent.C{"status", "=", "pedding"},
				},
				&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
			).
			Build()
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("should have a err when pass not accepted struct in P", func(t *testing.T) {
		_, _, err := goquent.New(goquent.MYSQL).
			Select("*").
			From("users").
			Where(
				&goquent.P{
					123,
					&goquent.C{"status", "=", "pedding"},
					"OR",
				},
				&goquent.C{"created_at", "BETWEEN", []string{"2023-01-01 00:00:00", "AND", "2023-12-31 23:59:59"}},
			).
			Build()
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("should return args in correct order", func(t *testing.T) {
		testArgs := []string{"pedding", "foo", "2023-01-01 00:00:00", "2023-12-31 23:59:59"}

		_, args, err := goquent.New(goquent.MYSQL).
			Select("*").
			From("users").
			Where(
				&goquent.P{
					&goquent.C{"status", "=", testArgs[0]},
					&goquent.C{"bar", "=", testArgs[1]},
				},
				&goquent.C{"created_at", "BETWEEN", []string{testArgs[2], "AND", testArgs[3]}},
			).
			Build()
		if err != nil {
			t.Error(err)
		}

		for i, v := range args {
			if v != testArgs[i] {
				t.Errorf("wanted %s gotted %s", testArgs[i], v)
			}
		}
	})

	t.Run("should optinal where in sql", func(t *testing.T) {
		filters := make([]goquent.Conditional, 0)

		sql, _, err := goquent.New(goquent.MYSQL).
			Select("*").
			From("users").
			Where(filters...).
			Build()
		if err != nil {
			t.Error(err)
		}

		testSql := "SELECT * FROM users"
		if strings.Trim(sql, " ") != testSql {
			t.Errorf("wanted %s gotted %s", testSql, sql)
		}
	})

	t.Run("should genereate pgsql query correct", func(t *testing.T) {
		sql, _, err := goquent.New(goquent.PGSQL).
			Select("*").
			From("users").
			Where(
				&goquent.C{"email", "=", "teste@email.com"},
				&goquent.C{"bar", "=", "foo"},
				&goquent.P{
					&goquent.C{"nested1", "=", "foo", "OR"},
					&goquent.C{"nested2", "=", "foo"},
				},
			).
			Build()
		if err != nil {
			t.Error(err)
		}

		testSql := "SELECT * FROM users WHERE email = $1 AND bar = $2 AND (nested1 = $3 OR nested2 = $4)"
		if strings.Trim(sql, " ") != testSql {
			t.Errorf("wanted %s gotted %s", testSql, sql)
		}
	})
}
