// https://www.geeksforgeeks.org/stack-set-4-evaluation-postfix-expression/

package main

import (
	"fmt"
	"os"
	"log"
)

type stack struct {
	size int
	operands [10]int
}

func (s *stack) Pop() int {
	if (s.Size()) <= 0 {
		log.Fatalln("Empty stack")
	}

	s.size--

	return s.operands[s.Size()]
}

func (s *stack) Push(opr int) {
	s.operands[s.Size()] = opr
	s.size++
}

func (s *stack) Size() int {
	return s.size
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<expression>")
		os.Exit(1)
	}

	var st stack
	operand := false

	for _, ch := range ([]byte(os.Args[1])) {
		if ch < '0' || ch > '9' {
			operand = false
		}

		switch {
			case ch == ' ':
			case ch >= '0' && ch <= '9':
				opr := int(ch - '0')
				if (operand) {
					opr = st.Pop()*10 + opr
				}
				st.Push(opr)
				operand = true
			case ch == '*':
				right, left := st.Pop(), st.Pop()
				st.Push(left*right)
			case ch == '/':
				right, left := st.Pop(), st.Pop()
				st.Push(left/right)
			case ch == '+':
				right, left := st.Pop(), st.Pop()
				st.Push(left+right)
			case ch == '-':
				right, left := st.Pop(), st.Pop()
				st.Push(left-right)
			default:
				log.Fatalln("Invalid expression.")
		}
	}

	fmt.Println(os.Args[1], "=", st.Pop())
}
