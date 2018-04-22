package main

import (
	"fmt"

	"github.com/m-okeefe/go-ds/ds"
)

/*
Test valid parentheses
Acceptable chars: () {} []
*/

func main() {
	t1 := "(hello!)"
	fmt.Println(isValidParen(t1))

	t2 := ")noway()in()correct"
	fmt.Println(isValidParen(t2))

	t3 := "it'(s (compl(c)at(ed).))"
	fmt.Println(isValidParen(t3))
}

// Iterative with stack, O(n) time, O(n) space
func isValidParen(in string) bool {

	s := ds.Stack{}

	for i := 0; i < len(in); i++ {
		if in[i] == '(' {
			s.Push(in[i])
		} else {
			if in[i] == ')' {
				if _, err := s.Peek(); err != nil {
					return false
				} else {
					s.Pop()
				}
			}
		}
	}

	// At the end, we should get a stack empty err; otherwise still have ( open paren still unclosed!
	if _, err := s.Peek(); err == nil {
		return false
	}
	return true
}
