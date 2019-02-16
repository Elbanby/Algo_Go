package trees

import "fmt"

type Tree struct {
	Head *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (t *Tree) Add(val int) {
	if t.Head == nil {
		t.Head = &Node{Value: val}
		return
	}
	t.Head.Add(val)
}

func (n *Node) Add(val int) {
	if n.Value <= val {
		if n.Right == nil {
			n.Right = &Node{Value: val}
			return
		}
		n.Right.Add(val)
		return
	}

	if n.Left == nil {
		n.Left = &Node{Value: val}
		return
	}
	n.Left.Add(val)
}

func (t *Tree) Remove(val int) {
	if t.Head == nil {
		return
	}
	t.Head = t.Head.Remove(val)
}

func (n *Node) Remove(val int) *Node {
	if n == nil {
		return n
	}

	if n.Value < val {
		n.Right = n.Right.Remove(val)
	} else if n.Value > val {
		n.Left = n.Left.Remove(val)
	} else {
		if n.Left == nil && n.Right == nil {
			n = nil
			return n
		}

		if n.Left == nil {
			n = n.Right
		} else if n.Right == nil {
			n = n.Left
		} else {
			tempNode := n.Right.FindMin()
			n.Value = tempNode.Value
			n.Right = n.Right.Remove(tempNode.Value)
		}
	}
	return n
}

func (n *Node) FindMin() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (t *Tree) InOrder() {
	if t.Head == nil {
		fmt.Println("Empty Tree")
		return
	}
	t.Head.InOrder()
	fmt.Println()
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%v ", n.Value)
	n.Right.InOrder()
}

func (t *Tree) PreOrder() {
	if t.Head == nil {
		fmt.Println("Empty Tree")
		return
	}
	t.Head.PreOrder()
	fmt.Println()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.Value)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (t *Tree) PostOrder() {
	if t == nil {
		return
	}
	t.Head.PostOrder()
	fmt.Println()
}

func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("%v ", n.Value)
}

func (t *Tree) LevelOrder() {
	if t.Head == nil {
		fmt.Println("Empty tree")
		return
	}
	t.Head.LevelOrder()
}

func (n *Node) LevelOrder() {
	// This will be a Queue slice implementation
	var q []*Node
	node := n
	q = append(q, node)

	for len(q) != 0 {
		fmt.Printf("%v ", q[0].Value)

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}

		if len(q) == 1 {
			break
		}

		q[0] = nil
		q = q[1:]
		node = q[0]
	}
	fmt.Println()
}
