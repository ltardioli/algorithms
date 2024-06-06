package parentesismatching

import "testing"

func TestSimpleBalance(t *testing.T) {
	t.Run("test simple balance. test case 1", func(t *testing.T) {
		exp := "(a + b) * (c + d)"
		if !CheckSimpleBalanced(exp) {
			t.Fatal("Expression is balanced but returned not balanced")
		}
	})

	t.Run("test simple balance. test case 2", func(t *testing.T) {
		exp := "((a + b) * (c + d))"
		if !CheckSimpleBalanced(exp) {
			t.Fatal("Expression is balanced but returned not balanced")
		}
	})

	t.Run("test simple balance. test case 3", func(t *testing.T) {
		exp := "((a + b) * (c + d)"
		if CheckSimpleBalanced(exp) {
			t.Fatal("Expression is not balanced but returned balanced")
		}
	})

	t.Run("test simple balance. test case 4", func(t *testing.T) {
		exp := "(a + b) * (c + d))"
		if CheckSimpleBalanced(exp) {
			t.Fatal("Expression is not balanced but returned balanced")
		}
	})

	t.Run("test simple balance. test case 5", func(t *testing.T) {
		exp := "((a + b) * (c + d))"
		if !CheckSimpleBalanced(exp) {
			t.Fatal("Expression is balanced but returned not balanced")
		}
	})
}

func TestBalance(t *testing.T) {
	t.Run("test balance. test case 1", func(t *testing.T) {
		exp := "{(a + b)} * [(c + d)]"
		if !CheckBalanced(exp) {
			t.Fatal("Expression is balanced but returned not balanced")
		}
	})
	t.Run("test balance. test case 2", func(t *testing.T) {
		exp := "{(a + b)} * [(c + d)]"
		if !CheckBalanced(exp) {
			t.Fatal("Expression is balanced but returned not balanced")
		}
	})

	t.Run("test balance. test case 3", func(t *testing.T) {
		exp := "{(a + b)} * [(c + d)}"
		if CheckBalanced(exp) {
			t.Fatal("Expression is not balanced but returned balanced")
		}
	})

	t.Run("test balance. test case 4", func(t *testing.T) {
		exp := "{(a + b)} * [(c + d)"
		if CheckBalanced(exp) {
			t.Fatal("Expression is not balanced but returned balanced")
		}
	})

	t.Run("test balance. test case 5", func(t *testing.T) {
		exp := "{(a + b)} * [(c + d)]}"
		if CheckBalanced(exp) {
			t.Fatal("Expression is not balanced but returned balanced")
		}
	})
}
