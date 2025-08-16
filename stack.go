package LinearDataStructures

import (
	"errors"
	"fmt"
)

type Stack struct {
	arr  []int
	top  int
	size int
}

// Constructor with preallocated slice
func NewStack(size int) *Stack {
	return &Stack{
		arr:  make([]int, 0, size),
		top:  -1,
		size: size,
	}
}

// Push adds an element to stack
func (s *Stack) Push(val int) {
	if s.top == s.size-1 {
		panic("stack is full")
	}
	s.arr = append(s.arr, val)
	s.top++
}

// Pop removes and returns the top element
func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return -1, errors.New("stack is empty")
	}
	topElement := s.arr[s.top]
	s.arr = s.arr[:s.top] 
	s.top--
	return topElement, nil
}

// Peek returns top element without removing
func (s *Stack) Peek() (int, error) {
	if s.top == -1 {
		return -1, errors.New("stack is empty")
	}
	return s.arr[s.top], nil
}

// Print displays the stack
func (s *Stack) Print() {
	if s.top == -1 {
		fmt.Println("stack is empty")
		return
	}
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.arr[i])
	}
}
