package linkedlist

import "github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"

type LinkedListInterface interface {
	Any() bool
	FindWithPrint(any) (node.NodeInterface, bool)
	Find(any) (node.NodeInterface, bool)
	First() node.NodeInterface
	Last() node.NodeInterface
	Insert(any, int) node.NodeInterface
	Pop() node.NodeInterface
	Print()
	Push(any) node.NodeInterface
	Shift() node.NodeInterface
	Unshift(any) node.NodeInterface
}
