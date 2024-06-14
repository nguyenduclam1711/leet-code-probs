package evaluatereversepolishnotation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		if t != "/" && t != "+" && t != "-" && t != "*" {
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
		} else {
			operand1, operand2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			if t == "/" {
				stack = append(stack, operand1/operand2)
			} else if t == "*" {
				stack = append(stack, operand1*operand2)
			} else if t == "+" {
				stack = append(stack, operand1+operand2)
			} else {
				stack = append(stack, operand1-operand2)
			}
		}
	}
	return stack[0]
}
