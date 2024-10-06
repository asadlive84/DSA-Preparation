package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	Top    *Node
	Bottom *Node
	Length int
}

func (s *Stack) push(data string) {
	node := &Node{
		Data: data,
	}
	// O(1)
	if s.Top == nil {
		s.Top = node
		s.Bottom = node
	} else {
		node.Next = s.Top
		s.Top = node
	}

	s.Length++
}

func (s *Stack) peek() string {
	// O(1)
	if s.Top == nil {
		return "Stack is empty"
	}
	return s.Top.Data
}

func (s *Stack) pop() {
	// O(1)
	if s.Top == nil {
		return
	}

	if s.Top.Next != nil {
		s.Top = s.Top.Next
		s.Length--
		return
	}
	s.Top = nil
	s.Bottom = nil
	s.Length = 0
}

func newStack(data string) *Stack {
	node := &Node{
		Data: data,
	}
	return &Stack{
		Top:    node,
		Bottom: node,
		Length: 1,
	}
}

func (s *Stack) show() {

	tmpTopNode := s.Top
	for tmpTopNode != nil {
		fmt.Printf("%+v \n", tmpTopNode.Data)
		tmpTopNode = tmpTopNode.Next
	}
}

func main() {
	stack := newStack("google")

	stack.push("yahoo")
	stack.push("udemey")
	stack.push("youtube")

	stack.show()

	fmt.Println("--------pop--------")
	// stack.pop()

	fmt.Println("--------peek--------")
	fmt.Printf("peek %+v\n", stack.peek())

	// stack.show()
}
