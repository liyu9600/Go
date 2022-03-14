package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {

	symbol := map[string]int{
		"+": -1,
		"-": -1,
		"*": -1,
		"/": -1,
	}

	stack := make([]int, 0)

	for _, v := range tokens {

		if symbol[v] == -1 {
			l := len(stack)
			a, b := stack[l-2], stack[l-1]
			r := count(a, b, v)
			if l >= 2 {
				stack = stack[0 : l-2]
			} else {
				stack = make([]int, 0)
			}

			stack = append(stack, r)
		} else {
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		}
	}

	return stack[0]
}

func count(a, b int, symbol string) int {

	switch symbol {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return 0
}
