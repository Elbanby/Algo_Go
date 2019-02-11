package stack

import "fmt"

type Stack struct {
	Top *Node
}

type Node struct {
	Value string
	Next  *Node
}

func newNode(val string) *Node {
	return &Node{Value: val}
}

func (s *Stack) Push(val string) {
	node := newNode(val)
	if s.Top == nil {
		s.Top = node
		return
	}
	s.Top.Push(node)
}

func (n *Node) Push(node *Node) {
	if n.Next == nil {
		n.Next = node
		return
	}
	n.Next.Push(node)
}

func (s *Stack) Pop() (string, bool) {
	if s.Top == nil {
		return "Nothing to pop", false
	}

	if s.Top.Next == nil {
		value := s.Top.Value
		s.Top = nil
		return value, true
	}
	return s.Top.Pop()
}

func (n *Node) Pop() (string, bool) {
	value := n.Value
	tail := n.Next

	if tail.Next == nil {
		value = n.Next.Value
		n.Next = nil
		return value, true
	}
	return tail.Pop()
}

func (s *Stack) Contains(val string) bool {
	if s.Top == nil {
		return false
	}
	node := s.Top
	for node != nil {
		if node.Value == val {
			return true
		}
		node = node.Next
	}
	return false
}

func (s *Stack) Print() {
	if s.Top == nil {
		fmt.Println("Stack is empty")
		return
	}
	node := s.Top
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func (s *Stack) Peek() {
	if s.Top == nil {
		fmt.Println("Stack is empty")
		return
	}
	node := s.Top
	for node.Next != nil {
		node = node.Next
	}
	fmt.Println(node.Value)
	return
}
