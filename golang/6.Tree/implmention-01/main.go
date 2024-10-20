package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newTree(data int) *Node {
	return &Node{
		Data: data,
	}
}

func (n *Node) insert(data int) {
	if data < n.Data {
		if n.Left == nil {
			n.Left = newTree(data)
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = newTree(data)
		} else {
			n.Right.insert(data)
		}
	}
}

func (n *Node) showData() {
	if n == nil {
		return
	}

	n.Left.showData()
	fmt.Printf("%+v\n", n.Data)
	n.Right.showData()
}

func (n *Node) showDataRecursive() {
	if n == nil {
		return
	}

	stack := []*Node{}
	current := n
	for current != nil || len(stack) > 0 {

		for current != nil {

			stack = append(stack, current)
			current = current.Left

		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("value: %+v ", current.Data)

		current = current.Right

	}
}

func main() {
	root := newTree(5)
	root.insert(4)
	root.insert(10)
	root.insert(25)
	root.insert(6)
	root.insert(3)
	root.insert(2)

	// root.showData()

	fmt.Println("===========")
	root.showDataRecursive()
}
