package main

import "fmt"

type UniqueStack struct {
	data map[string]struct{}
	elements []string
}

func NewUniqueStack() *UniqueStack {
	return &UniqueStack{
		data: make(map[string]struct{}),
		elements: []string{},
	}
}

func (s *UniqueStack) push(value string) {
	if _, exists := s.data[value]; !exists {
		s.data[value] = struct{}{}
		s.elements = append(s.elements, value)
	}
}

func (s *UniqueStack) pop() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	delete(s.data, value)
	return value, true
}

func (s *UniqueStack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *UniqueStack) value() []string {
	return s.elements
}

func main() {
	stack := NewUniqueStack()
	fmt.Println(stack.isEmpty())
	stack.push("civic")
	stack.push("avanza")
	stack.push("brio")
	stack.push("lamborghini")

	fmt.Println(stack.value())

	value, ok := stack.pop()

	fmt.Println(value, ok)
	fmt.Println(stack.value())
	fmt.Println(stack.isEmpty())
}
