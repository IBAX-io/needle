package vm

import (
	"github.com/pkg/errors"
)

// ErrStackUnderFlow is an error that is returned when the stack is underflow.
var ErrStackUnderFlow = errors.New("stack under flow")

// Stack is a stack that contains elements of any type.
type Stack struct {
	element []any
}

// newStack creates a new stack with an initial capacity of 1024.
func newStack() *Stack {
	return &Stack{element: make([]any, 0, 1024)}
}

// Stack returns the stack as a slice.
func (s *Stack) Stack() []any {
	return s.element
}

// size returns the current size of the stack.
func (s *Stack) size() int {
	return len(s.element)
}

// CheckDepth checks if the stack has at least min elements.
func (s *Stack) CheckDepth(min int) error {
	if s.size() < min {
		return ErrStackUnderFlow
	}
	return nil
}

// set sets the element at the given index in the stack.
func (s *Stack) set(ind int, d any) {
	s.element[ind] = d
}

// reset clears all elements from the stack.
func (s *Stack) reset() {
	s.element = s.element[:0]
}

// push adds an element to the top of the stack.
func (s *Stack) push(d any) {
	s.element = append(s.element, d)
}

// pushN adds multiple elements to the top of the stack.
func (s *Stack) pushN(d []any) {
	s.element = append(s.element, d...)
}

// get returns the element at the given index from the stack.
func (s *Stack) get(idx int) any {
	if idx >= 0 && s.size() > 0 && s.size() > idx {
		return s.element[idx]
	}
	return nil
}

// getAndDel returns and removes the element at the given index from the stack.
func (s *Stack) getAndDel(idx int) any {
	if idx >= 0 && s.size() > 0 && s.size() > idx {
		ret := s.element[idx]
		s.element = append(s.element[:idx], s.element[idx+1:]...)
		return ret
	}
	return nil
}

// peek returns the top element from the stack.
func (s *Stack) peek() any {
	if s.size() == 0 {
		return nil
	}
	return s.element[s.size()-1]
}

// peekN returns the top n elements from the stack.
func (s *Stack) peekN(n int) []any {
	sLen := len(s.element)
	var el []any = nil
	if sLen >= n {
		el = s.element[sLen-n:]
	}
	ret := make([]any, n)
	copy(ret, el)
	return ret
}

// peekFromTo returns elements from index from to index to.
func (s *Stack) peekFromTo(from, to int) []any {
	var el []any = nil
	if from >= 0 && to >= 0 && from <= to && s.size() >= to {
		el = s.element[from:to]
	}
	return el
}

// pop returns and removes the top element from the stack.
func (s *Stack) pop() (ret any) {
	if s.size() == 0 {
		s.element = s.element[:0]
		return nil
	}
	ret = s.element[s.size()-1]
	s.element = s.element[:s.size()-1]
	return
}

// popN returns and removes the top n elements from the stack.
func (s *Stack) popN(n int) []any {
	sLen := len(s.element)
	if sLen < n {
		n = sLen
	}
	elem := s.element[sLen-n:]
	s.element = s.element[:sLen-n]
	ret := make([]any, n)
	copy(ret, elem)
	return ret
}

// swap swaps the nth element from the top with the top element of the stack.
func (s *Stack) swap(n int) {
	s.element[s.size()-n], s.element[s.size()-1] = s.peek(), s.element[s.size()-n]
}

// dup duplicates the nth element from the top of the stack.
func (s *Stack) dup(n int) {
	s.push(&s.element[s.size()-n])
}

// resetByIdx removes all elements from the given index to the top of the stack.
func (s *Stack) resetByIdx(idx int) {
	if idx < 0 || idx > s.size() {
		idx = 0
	}
	s.element = s.element[:idx]
}
