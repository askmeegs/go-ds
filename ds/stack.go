package ds

import "fmt"

const ErrStackEmpty = "Stack is empty"
const ErrNilItem = "Item to push is nil"

// slice-based stack
// space complexity: O(n)
type Stack struct {
	slice []interface{} // last element is the "top"
}

// O(1)
func (s *Stack) Push(item interface{}) error {
	if item == nil {
		return fmt.Errorf(ErrNilItem)
	}
	s.slice = append(s.slice, item)
	return nil
}

// O(1)
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf(ErrStackEmpty)
	}
	rem := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return rem, nil
}

// O(1)
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf(ErrStackEmpty)
	}
	return s.slice[len(s.slice)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.slice) == 0 {
		return true
	}
	return false
}

func (s *Stack) ToString() string {
	output := ""
	for i := len(s.slice) - 1; i >= 0; i-- {
		output = fmt.Sprintf("%s\n[ %v ]", output, s.slice[i])
	}
	return output
}
