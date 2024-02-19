package main

import "fmt"

type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}
	length := len(*s) - 1
	delElement := (*s)[length]
	*s = (*s)[:length]
	fmt.Println("The deleted element is", delElement)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func main() {
	var data Stack
	data.push(1)
	data.push(2)
	data.push(3)
	data.push(4)
	data.push(5)
	data.push(6)

	data.pop()

	fmt.Println(data)
}
