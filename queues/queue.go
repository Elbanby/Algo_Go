package queue

import "fmt"

type Queue struct {
	Head *Node
}

type Node struct {
	Value string
	Next  *Node
}

func (q *Queue) Enqueue(val string) {
	if q.Head == nil {
		q.Head = &Node{Value: val}
		return
	}
	q.Head.Enqueue(val)
}

func (n *Node) Enqueue(val string) {
	if n.Next == nil {
		n.Next = &Node{Value: val}
		return
	}
	n.Next.Enqueue(val)
}

func (q *Queue) Deque() (string, bool) {
	if q.Head == nil {
		return "Queue was empty hence can't deque", false
	}

	if q.Head.Next != nil {
		value := q.Head.Value
		q.Head = q.Head.Next
		return fmt.Sprintf("Dequed value: %v", value), true
	}
	q.Head = nil
	return "Queue is erased", true
}

func (q *Queue) Print() {
	if q.Head == nil {
		fmt.Println("Queue is currently empty")
		return
	}
	node := q.Head
	for node != nil {
		fmt.Printf("%v ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func (q *Queue) Contain(val string) bool {
	if q.Head == nil {
		fmt.Println("Queue is empty\n")
		return false
	}
	node := q.Head
	for node != nil {
		if node.Value == val {
			return true
		}
		node = node.Next
	}
	return false
}
