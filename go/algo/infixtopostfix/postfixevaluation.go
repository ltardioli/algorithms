package infixtopostfix

import "algostudy/utils"

// It will only work for single digit numbers
func Evaluate(expression string) int {
	var stack utils.Stack[int]
	for _, c := range expression {
		if !isOperator(c) {
			stack.Push(int(c - '0'))
		} else {
			switch c {
			case '+':
				y, _ := stack.Pop()
				x, _ := stack.Pop()
				r := x + y
				stack.Push(r)
			case '-':
				y, _ := stack.Pop()
				x, _ := stack.Pop()
				r := x - y
				stack.Push(r)
			case '*':
				y, _ := stack.Pop()
				x, _ := stack.Pop()
				r := x * y
				stack.Push(r)
			case '/':
				y, _ := stack.Pop()
				x, _ := stack.Pop()
				r := x / y
				stack.Push(r)
			}
		}
	}
	r, _ := stack.Pop()
	return r
}
