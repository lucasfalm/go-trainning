package linkedlist

import "fmt"

func (l *LinkedList) Print() {
	tree := ""

	for _, node := range l.nodes {
		tree += fmt.Sprintf("%v <-> ", node.Value())
	}

	if len(tree) > 0 {
		tree = tree[:len(tree)-4]
	}

	fmt.Println(tree)
}

func (n *Node) Print() {
	tree := ""

	if n.Head() != nil {
		tree += fmt.Sprintf("%v <-> ", n.head.Value())
	} else {
		tree += fmt.Sprintf("%v <-> ", "nil")
	}

	tree += fmt.Sprintf("%v", n.value)

	if n.Tail() != nil {
		tree += fmt.Sprintf(" <-> %v", n.tail.Value())
	} else {
		tree += fmt.Sprintf(" <-> %v", "nil")
	}

	fmt.Println(tree)
}
