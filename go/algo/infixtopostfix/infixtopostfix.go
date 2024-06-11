package infixtopostfix

import "algostudy/utils"

type precedence map[rune]int

func createProcedenceTable() precedence {
	return precedence{'+': 1, '-': 1, '*': 2, '/': 2}
}

func hasKey(precTable precedence, key rune) bool {
	_, exists := precTable[key]
	return exists
}

func Convert(expression string) string {
	precTable := createProcedenceTable()
	var stack utils.Stack[rune]
	var postfix string

	for _, c := range expression {
		if !hasKey(precTable, c) {
			postfix += string(c)
		} else {
			check := true
			for check {
				top, _ := stack.Peek()
				if stack.IsEmpty() || precTable[c] > precTable[top] {
					stack.Push(c)
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
