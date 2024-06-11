package infixtopostfix

import (
	"testing"
)

func TestPostFixEvaluation(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		expr := "12*63/+"
		expect := 4
		result := Evaluate(expr)

		if result != expect {
			t.Fatalf("expected %v, found %v", expect, result)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		expr := "52+"
		expect := 7
		result := Evaluate(expr)
		if result != expect {
			t.Fatalf("expected %v, found %v", expect, result)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		expr := "83-"
		expect := 5
		result := Evaluate(expr)
		if result != expect {
			t.Fatalf("expected %v, found %v", expect, result)
		}
	})

	t.Run("test case 4", func(t *testing.T) {
		expr := "42*31*+"
		expect := 11
		result := Evaluate(expr)
		if result != expect {
			t.Fatalf("expected %v, found %v", expect, result)
		}
	})

}
