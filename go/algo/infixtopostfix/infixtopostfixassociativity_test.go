package infixtopostfix

import "testing"

func TestInfixToPostFixWithAssocistivityConversion(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expr := "((a+b)*c)-d^e^f"
		converted := ConvertAssociativity(expr)
		expect := "ab+c*def^^-"

		if converted != expect {
			t.Fatalf("expected %v, found %v", expect, converted)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		expr := "a+b*c-d/e^f"
		converted := ConvertAssociativity(expr)
		expect := "abc*+def^/-"

		if converted != expect {
			t.Fatalf("expected %v, found %v", expect, converted)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		expr := "a+b*c/d^e-f"
		converted := ConvertAssociativity(expr)
		expect := "abc*de^/+f-"

		if converted != expect {
			t.Fatalf("expected %v, found %v", expect, converted)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		expr := "a+b*c-d^e/f"
		converted := ConvertAssociativity(expr)
		expect := "abc*+de^f/-"

		if converted != expect {
			t.Fatalf("expected %v, found %v", expect, converted)
		}
	})
}
