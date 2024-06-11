package infixtopostfix

import "testing"

func TestInfixToPostFixConversion(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expr := "a+b+c"
		expect := "ab+c+"
		converted := Convert(expr)
		if converted != expect {
			t.Fatalf("expected %s, found %s", expect, converted)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expr := "a+b*c-d/e"
		expect := "abc*+de/-"
		converted := Convert(expr)
		if converted != expect {
			t.Fatalf("expected %s, found %s", expect, converted)
		}
	})
}
