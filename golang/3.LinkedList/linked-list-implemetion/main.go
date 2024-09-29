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

func newLinkedList(data int) LinkedList {
	n := &Node{
		Data: data,
	}
	return LinkedList{
		Head:   n,
		Tail:   n,
		Length: 1,
	}
}

func (ll *LinkedList) Prepend(existsData, data int) {
	node := &Node{
		Data: data,
	}

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		ll.Length = 1
		return
	}

	if ll.Head.Data == existsData {
		node.Next = ll.Head
		ll.Head = node
		ll.Length++
		return
	}

	currentNode := ll.Head
	nextNode := currentNode.Next

	// newNode

	for nextNode != nil {
		if nextNode.Data == existsData {
			node.Next = nextNode
			currentNode.Next = node
			ll.Length++
			return
		}
		currentNode = currentNode.Next
		nextNode = currentNode.Next
	}
}

func (ll *LinkedList) Insert(data int) {
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

func (ll *LinkedList) show() []int {
	tempNode := ll.Head
	llArray := make([]int, 0)
	for tempNode != nil {
		llArray = append(llArray, tempNode.Data)
		tempNode = tempNode.Next
	}
	return llArray
}

func main() {

	ll := newLinkedList(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	fmt.Println("================================== ======")
	fmt.Printf("%+v\n", ll.show())
	fmt.Printf("%+v\n", ll.Length)
	fmt.Println("========================================")

	ll.Prepend(5, 10)
	// ll.Prepend(2, 11)

	fmt.Println("================================== ======")
	fmt.Printf("%+v\n", ll.show())
	fmt.Printf("%+v\n", ll.Length)
	fmt.Println("========================================")

}
