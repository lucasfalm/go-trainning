package linkedlist

import "fmt"

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
