package goquent_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lucasres/goquent"
)

func TestWhereInstruction(t *testing.T) {
	t.Run("should create a where clause", func(t *testing.T) {
		sql, args, err := goquent.New().Where(goquent.C{"status", "=", "actived"}).Build()
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
		sql, args, err := goquent.New().Where(
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
		sql, args, err := goquent.New().Where(
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

		_, args, err := goquent.New().Where(
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
		_, _, err := goquent.New().Where(
			goquent.C{"status", "=", "actived", "OR", "bar", "foo"},
		).Build()
		if err == nil {
			t.Error(errors.New("invalid args must return an error"))
		}

		_, _, err = goquent.New().Where(
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

		sql, args, err := goquent.New().Where(
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
