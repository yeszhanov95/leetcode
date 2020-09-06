package main

import "strconv"

func evalRPN(tokens []string) int {
	operands := make([]int, 0, len(tokens))

	for _, str := range tokens {
		if num, err := strconv.Atoi(str); err == nil {
			operands = append(operands, num)
			continue
		}

		leftOperand, rightOperand := operands[len(operands)-2], operands[len(operands)-1]
		operands = operands[:len(operands)-2]

		var resultOperand int
		switch str {
		case "+": resultOperand = leftOperand + rightOperand
		case "-": resultOperand = leftOperand - rightOperand
		case "*": resultOperand = leftOperand * rightOperand
		case "/": resultOperand = leftOperand / rightOperand
		}

		operands = append(operands, resultOperand)
	}

	return operands[0]
}