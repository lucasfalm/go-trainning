package linkedlist

type LinkedListInterface interface {
	Any() bool
	FindWithPrint(any) (NodeInterface, bool)
	Find(any) (NodeInterface, bool)
	First() NodeInterface
	Last() NodeInterface
	Insert(any, int) NodeInterface
	Pop() NodeInterface
	Print()
	Push(any) NodeInterface
	Shift() NodeInterface
	Unshift(any) NodeInterface
}

// NOTE: doubly linked list
type LinkedList struct {
	nodes []NodeInterface
	count int
}

// NOTE: creates a linked list and returns the interface/contract
func NewLinkedList() LinkedListInterface {
	return &LinkedList{}
}
