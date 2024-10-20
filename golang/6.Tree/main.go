package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

func (t *Node) insert(data int) {
	if data <= t.Data {
		if t.Left == nil {
			t.Left = &Node{
				Data: data,
			}
		} else {
			t.Left.insert(data)
		}
	} else {
		if t.Right == nil {
			t.Right = &Node{
				Data: data,
			}
		} else {
			t.Right.insert(data)
		}
	}
}

func (t *Node) inOrderShow() {
	if t == nil {
		return
	}
	t.Left.inOrderShow()
	fmt.Printf("Data: %+v\n", t.Data)
	t.Right.inOrderShow()

}

func (t *Node) postOrderShow() {
	if t == nil {
		return
	}
	fmt.Printf("%+v\n", t.Data)
	t.Left.inOrderShow()
	t.Right.inOrderShow()
}

func (t *Node) lookup(data int) bool {

	if t == nil {
		return false
	}

	if t.Data == data {
		return true
	} else if data < t.Data {
		return t.Left.lookup(data)
	} else {
		return t.Right.lookup(data)
	}
}

func (t *Node) remove(data int) *Node {
	if t == nil {
		return nil
	}
	if data < t.Data {
		return t.Left.remove(data)
	} else if data > t.Data {
		return t.Right.remove(data)
	} else {
		if t.Right == nil && t.Left == nil {
			return nil
		}

		if t.Right == nil {
			return t.Left
		}

		if t.Left == nil {
			return t.Right
		}

		successor := findMin(t.Right)
		t.Data = successor.Data
		t.Right = t.Right.remove(successor.Data)
	}

	return t

}

func findMin(node *Node) *Node {
	currentNode := node

	for currentNode != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func main() {
	root := newNode(75)

	root.insert(85)
	root.insert(65)
	root.insert(95)
	root.insert(60)
	root.insert(70)

	// root.inOrderShow()
	root.postOrderShow()

	// fmt.Printf("Lookup: %+v\n", root.lookup(71))

}
