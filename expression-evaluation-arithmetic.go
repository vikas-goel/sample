// https://www.geeksforgeeks.org/arithmetic-expression-evalution/

package main

import (
	"fmt"
	"log"
	"os"
)

type Operand struct {
	size int
	elems [10]int
}

func (o *Operand) Size() int {
	return o.size
}

func (o *Operand) Pop() int {
	if o.Size() < 1 {
		log.Fatalln("Invalid Operand::Pop() call.")
	}

	o.size--
	return o.elems[o.Size()]
}

func (o *Operand) Push(i int) {
	o.elems[o.Size()] = i
	o.size++
}

type Operator struct {
	size int
	elems [10]byte
}

func (o *Operator) Size() int {
	return o.size
}

func (o *Operator) Peek() byte {
	if o.Size() < 1 {
		log.Fatalln("Invalid Operator::Peek() call.")
	}

	return o.elems[o.Size()-1]
}

func (o *Operator) Pop() byte {
	if o.Size() < 1 {
		log.Fatalln("Invalid Operator::Pop() call.")
	}

	o.size--
	return o.elems[o.Size()]
}

func (o *Operator) Push(b byte) {
	o.elems[o.Size()] = b
	o.size++
}

func evaluate(oper byte, right, left int) (result int) {
	fmt.Printf("%c ", oper)
	switch oper {
		case '*':
			result = left * right
		case '/':
			result = left / right
		case '+':
			result = left + right
		case '-':
			result = left - right
		default:
			log.Fatalln("Invalid operator.")
	}

	return
}

func hasPrecedence(oper1, oper2 byte) bool {
	if oper2 == '(' || oper2 == ')' {
		return false
	}

	if (oper1 == '*' || oper1 == '/') && (oper2 == '+' || oper2 == '-') {
		return false
	}

	return true
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<expression>")
	}

	var operands Operand
	var operator Operator

	operOn := false
	expr := []byte(os.Args[1])
	for _, ch := range expr {
		if ch < '0' || ch > '9' {
			operOn = false
		}

		switch {
			case ch == ' ':
			case ch >= '0' && ch <= '9':
				opr := int(ch - '0')
				if operOn {
					opr = operands.Pop()*10 + opr
				}
				operands.Push(opr)
				operOn = true
			case ch == '(':
				operator.Push(ch)
			case ch == ')':
				for oper := operator.Pop(); oper != '('; oper = operator.Pop() {
					operands.Push(evaluate(oper, operands.Pop(), operands.Pop()))
				}
			case ch == '*' || ch == '/' || ch == '+' || ch == '-':
				for operator.Size() > 0 && hasPrecedence(ch, operator.Peek()) {
					operands.Push(evaluate(operator.Pop(), operands.Pop(), operands.Pop()))
				}
				operator.Push(ch)
			default:
				log.Fatalln("Invalid expression.")
		}
	}

	for operator.Size() > 0 {
		operands.Push(evaluate(operator.Pop(), operands.Pop(), operands.Pop()))
	}

	fmt.Println(os.Args[1], "=", operands.Pop())
}
