package parentesismatching

import "algostudy/utils"

func CheckSimpleBalanced(expression string) bool {
	var stack utils.Stack[rune]
	for _, s := range expression {
		if s == '(' {
			stack.Push(s)
		} else if s == ')' {
			_, error := stack.Pop()
			if error != nil {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func CheckBalanced(expression string) bool {
	var stack utils.Stack[rune]
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, s := range expression {
		if s == '(' || s == '[' || s == '{' {
			stack.Push(s)
		} else if s == ')' || s == ']' || s == '}' {
			peek, error := stack.Peek()
			if error != nil {
				return false
			}

			if matching[s] != peek {
				return false
			} else {
				_, error := stack.Pop()
				if error != nil {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}
