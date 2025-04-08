package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
func isValid(s string) bool {
	stack := Stack[string]{}
	for _, c := range s {
		switch c {
		case '{':
			stack.Push(string(c))
		case '(':
			stack.Push(string(c))
		case '[':
			stack.Push(string(c))
		case '}':
			top, _ := stack.Top()
			if top == "{" {
				stack.Pop()
			} else {
				return false
			}
		case ')':
			top, _ := stack.Top()
			if top == "(" {
				stack.Pop()
			} else {
				return false
			}
		case ']':
			top, _ := stack.Top()
			if top == "[" {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func main() {
	str := "{}[]()}"
	fmt.Println(isValid(str))
}
