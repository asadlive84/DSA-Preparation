package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Queue struct {
	Front *Node
	Back  *Node
	Size  int
}

func (q *Queue) enQueue(data string) {
	node := &Node{
		Data: data,
	}
	if q.Front == nil {
		q.Front = node
		q.Back = node
		q.Size = 1
		return
	}

	q.Back.Next = node
	q.Back = node
	q.Size++
}

func (q *Queue) deQueue() {
	if q.Front == nil {
		return
	}

	if q.Front.Next == nil {
		q.Front = nil
		q.Back = nil
	} else {
		q.Front = q.Front.Next
	}
	q.Size--
}

func (q *Queue) peek() string {
	if q.Front == nil {
		return "no data left"
	}
	return q.Front.Data
}

func (q *Queue) show() []string {
	if q.Front == nil {
		return []string{""}
	}
	tmpNode := q.Front

	nodeList := make([]string, 0, q.Size)

	for tmpNode != nil {
		nodeList = append(nodeList, tmpNode.Data)

		tmpNode = tmpNode.Next
	}
	return nodeList
}

func newQueue(data string) *Queue {
	node := &Node{
		Data: data,
	}
	return &Queue{
		Front: node,
		Back:  node,
		Size:  1,
	}
}

func main() {

	queue := newQueue("google")
	queue.enQueue("yahooo")
	queue.enQueue("youtube")
	queue.enQueue("fb")
	queue.enQueue("iig")
	queue.enQueue("iig")

	fmt.Printf("show data %+v\n", queue.show())
	fmt.Printf("show Size %+v\n", queue.Size)
	fmt.Println("===========")
	queue.deQueue()
	queue.deQueue()
	queue.deQueue()
	queue.deQueue()
	queue.deQueue()
	queue.deQueue()
	fmt.Printf("show peek %+v\n", queue.peek())
	fmt.Printf("show Size %+v\n", queue.Size)
	fmt.Println("===========")

	queue.enQueue("asad")
	fmt.Printf("show peek %+v\n", queue.peek())
	fmt.Printf("show Size %+v\n", queue.Size)

}
