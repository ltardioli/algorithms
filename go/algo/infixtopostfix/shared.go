package infixtopostfix

func isOperand(x rune) bool {
	if x == '+' || x == '-' || x == '*' || x == '/' ||
		x == '^' || x == '(' || x == ')' {
		return false
	}
	return true
}

func isOperator(x rune) bool {
	if x == '+' || x == '-' || x == '*' || x == '/' {
		return true
	}
	return false
}
