package linkedList

import "fmt"

type List struct {
	Head *Node
}

type Node struct {
	Value string
	Next  *Node
}

func (l *List) Add(val string) {
	if l.Head == nil {
		l.Head = &Node{Value: val}
		return
	}
	l.Head.Add(val)
}

func (n *Node) Add(val string) {
	if n.Next == nil {
		n.Next = &Node{Value: val}
		return
	}
	n.Next.Add(val)
}

func (l *List) Remove(val string) (string, bool) {
  if l.Head == nil {
    return "List was empty, nothing is deleted", false
  }

  head := l.Head
  node := l.Head.Next

  if head.Value == val  && head.Next == nil {
    l.Head = nil
    return "List deleted", true
  }

  for node != nil {
    if l.Head.Value == val {
      l.Head = l.Head.Next
      return fmt.Sprintf("Removed %v from list", val), true
    }

    if node.Value == val {
      head.Next = node.Next
      return fmt.Sprintf("Removed %v from list", val), true
    }
    head = node
    node = node.Next
  }
  return "Value not found", false
}

func (l *List) Contains(val string) bool {
  if l.Head == nil {
    return false
  }
  node := l.Head
  for node != nil {
    if node.Value == val {
      return true
    }
    node = node.Next
  }
  return false
}

func (l *List) Print() {
  if l.Head == nil {
    fmt.Println("List is empty! nothing to print")
    return
  }
  node := l.Head
  for node != nil {
    fmt.Printf("%v -> ", node.Value)
    node = node.Next
  }
  fmt.Println("(/)")
}
