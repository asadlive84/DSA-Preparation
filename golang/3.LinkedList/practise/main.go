package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func newLinkedList(data int) *LinkedList {
	node := &Node{
		Data: data,
		Next: nil,
	}
	return &LinkedList{
		Head:   node,
		Tail:   node,
		Length: 1,
	}
}

func (ll *LinkedList) append(data int) {
	node := &Node{
		Data: data,
	}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		ll.Length = 1
		return
	}

	ll.Tail.Next = node
	ll.Tail = node
	ll.Length++

}

func (ll *LinkedList) prepend(data int) {
	node := &Node{
		Data: data,
	}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		ll.Length = 1
		return
	}

	node.Next = ll.Head
	ll.Head = node
	ll.Length++
}

func (ll *LinkedList) insertByIndex(index, data int) {
	node := &Node{
		Data: data,
	}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		ll.Length = 1
		return
	}

	if index == 0 {
		node.Next = ll.Head
		ll.Head = node
		if ll.Length == 1 {
			ll.Tail = node
		}
		ll.Length++
		return
	}

	currentNode := ll.Head

	for i := 0; currentNode != nil; i++ {
		if index == i {
			node.Next = currentNode.Next
			currentNode.Next = node
			if node.Next == nil {
				ll.Tail = node
			}
			ll.Length++
			return
		}

		currentNode = currentNode.Next
	}

}

func (ll *LinkedList) deleteByIndex(index int) {

	if ll.Head == nil {
		return
	}

	if index == 0 {
		tmpNode := ll.Head.Next
		ll.Head = tmpNode
		if tmpNode == ll.Tail {
			ll.Tail = tmpNode
		}
		ll.Length--
		return
	}

	currentNode := ll.Head
	for i := 1; i < index; i++ {
		currentNode = currentNode.Next
	}

	if currentNode.Next == ll.Tail {
		ll.Tail = currentNode
	}
	currentNode.Next = currentNode.Next.Next
	ll.Length--

}
func (ll *LinkedList) show() []int {
	list := make([]int, 0, ll.Length)

	tmpNode := ll.Head

	for tmpNode != nil {
		list = append(list, tmpNode.Data)
		tmpNode = tmpNode.Next

	}

	return list
}

func (ll *LinkedList) appendMultiple(ds ...int) {
	for _, k := range ds {
		ll.append(k)
	}
}

func (ll *LinkedList) hasCycle() bool {
	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func (ll *LinkedList) reverse() {

	var prev *Node
	var current = ll.Head
	var next *Node
	ll.Tail = ll.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev

}

func main() {
	ll := newLinkedList(1)

	ll.append(2)
	ll.append(3)
	ll.append(4)

	fmt.Println("========================================")
	fmt.Printf("data: %+v\n", ll.show())
	fmt.Println("========================================")

	ll.insertByIndex(2, 99)
	ll.insertByIndex(0, 78)
	ll.insertByIndex(4, 798)
	// ll.prepend(100)

	fmt.Println("========================================")
	fmt.Printf("data: %+v\n", ll.show())
	fmt.Println("========================================")

	fmt.Printf("length: %+v\n", ll.Length)
	// ll.deleteByIndex(0)
	// ll.deleteByIndex(2)
	// ll.deleteByIndex(2)
	// ll.deleteByIndex(2)
	// ll.deleteByIndex(2)

	// fmt.Println("========================================")
	// fmt.Printf("data: %+v\n", ll.show())
	// fmt.Println("========================================")
	// // ll.deleteByIndex(2)
	// // ll.deleteByIndex(0)
	// // ll.deleteByIndex(1)
	// // ll.deleteByIndex(2)
	// ll.deleteByIndex(0)
	ll.deleteByIndex(6)
	ll.deleteByIndex(0)

	fmt.Println("========================================")
	fmt.Printf("length: %+v\n", ll.Length)
	fmt.Printf("data: %+v\n", ll.show())
	fmt.Printf("hasCycle(): %+v\n", ll.hasCycle())
	fmt.Println("========================================")

	ll.reverse()

	fmt.Println("==============##reverse##==========================")
	fmt.Printf("length: %+v\n", ll.Length)
	fmt.Printf("data: %+v\n", ll.show())
	fmt.Printf("data: %+v\n", ll.Head)
	fmt.Printf("data: %+v\n", ll.Tail)
	fmt.Println("========================================")

}
