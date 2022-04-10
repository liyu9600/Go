package main

func longestValidParentheses(s string) int {

	maxLength := 0
	stack := []int{-1}

	for i, v := range s {

		if v == '(' {
			stack = append(stack, i)
			continue
		}

		l := len(stack)
		if l > 1 {
			stack = stack[0:l-1]
			if r := i - stack[l-2]; r > maxLength {
				maxLength = r
			}

			continue
		}

		stack[0] = i
	}

	return maxLength
}
