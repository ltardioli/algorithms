package infixtopostfix

import "algostudy/utils"

type inPrecedence map[rune]int
type outPrecedence map[rune]int

func createProcedenceTables() (inPrecedence, outPrecedence) {
	return inPrecedence{'+': 2, '-': 2, '*': 4, '/': 4, '^': 5, '(': 0}, outPrecedence{'+': 1, '-': 1, '*': 3, '/': 3, '^': 6, '(': 7, ')': 0}
}

func ConvertAssociativity(expression string) string {
	var stack utils.Stack[rune]
	var postfix string
	inPrecedence, outPrecedence := createProcedenceTables()

	for _, c := range expression {
		if isOperand(c) {
			postfix += string(c)
		} else {
			check := true
			for check {
				top, _ := stack.Peek()
				if stack.IsEmpty() || outPrecedence[c] > inPrecedence[top] {
					stack.Push(c)
					check = false
				} else if outPrecedence[c] == inPrecedence[top] {
					stack.Pop()
					check = false
				} else {
					pop, _ := stack.Pop()
					postfix += string(pop)
				}
			}
		}
	}

	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		postfix += string(pop)
	}
	return postfix

}
