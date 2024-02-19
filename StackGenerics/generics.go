package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) pop() {
	if s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}
	length := len(*s) - 1
	delElement := (*s)[length]
	*s = (*s)[:length]
	fmt.Println("The deleted element is", delElement)
}

func (s *Stack[T]) isEmpty() bool {
	return len(*s) == 0
}
func main() {
	stack := Stack[int]{}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()

	fmt.Println(stack)
}
