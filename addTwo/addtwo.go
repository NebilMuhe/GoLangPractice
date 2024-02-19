package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func addTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	x := ""
	y := ""
	for l1 != nil || l2 != nil {
		x += strconv.Itoa(l1.val)
		y += strconv.Itoa(l2.val)
		l1 = l1.Next
		l2 = l2.Next
	}

	first, _ := strconv.Atoi(x)
	second, _ := strconv.Atoi(y)
	sum := first + second

	l3 := &ListNode{}
	str := strconv.Itoa(sum)
	tail := l3
	for i := len(str) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Println("Something went wrong")
		}
		tail.val = digit
		node := &ListNode{}
		if i == 0 {
			tail.Next = nil
		} else {
			tail.Next = node
			tail = node
		}
	}
	return l3
}

func main() {
	x := &ListNode{val: 2, Next: &ListNode{val: 4, Next: &ListNode{val: 3, Next: nil}}}
	y := &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4, Next: nil}}}

	val := addTwo(x, y)

	for val != nil {
		fmt.Println(val.val)
		val = val.Next
	}
}
