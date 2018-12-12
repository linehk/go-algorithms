// Package dijkstraDoubleStack implements math expression by Dijkstra double stack.
package dijkstraDoubleStack

import (
	"GoAlgorithms/DataStructures/stack/linkedListStack"
)

func Eval(expression string) int {
	operatorStack := linkedListStack.New()
	valueStack := linkedListStack.New()

	for _, ch := range expression {
		switch ch {
		case '(':
		case '+':
			operatorStack.Push(ch)
		case '-':
			operatorStack.Push(ch)
		case '*':
			operatorStack.Push(ch)
		case '/':
			operatorStack.Push(ch)
		case ')':
			opInterface, _ := operatorStack.Pop()
			op, _ := opInterface.(int32)
			valInterface, _ := valueStack.Pop()
			firstVal, _ := valInterface.(int)
			firstVal -= '0'
			switch op {
			case '+':
				valInterface, _ := valueStack.Pop()
				secondVal, _ := valInterface.(int)
				secondVal -= '0'
				firstVal += secondVal
			case '-':
				valInterface, _ := valueStack.Pop()
				secondVal, _ := valInterface.(int)
				secondVal -= '0'
				firstVal -= secondVal
			case '*':
				valInterface, _ := valueStack.Pop()
				secondVal, _ := valInterface.(int)
				secondVal -= '0'
				firstVal *= secondVal
			case '/':
				valInterface, _ := valueStack.Pop()
				secondVal, _ := valInterface.(int)
				secondVal -= '0'
				firstVal /= secondVal
			}
			valueStack.Push(firstVal)
		default:
			valueStack.Push(int(ch))
		}
	}
	valInterface, _ := valueStack.Pop()
	v, _ := valInterface.(int)
	return v
}
