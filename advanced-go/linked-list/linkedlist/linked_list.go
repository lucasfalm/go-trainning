package linkedlist

type LinkedListInterface interface {
	Any() bool
	FindWithPrint(any) (*Node, bool)
	Find(any) (*Node, bool)
	First() *Node
	Last() *Node
	Insert(any, int) *Node
	Pop() *Node
	Print()
	Push(any) *Node
	Shift() *Node
	Unshift(any) *Node
}

// NOTE: doubly linked list
type LinkedList struct {
	nodes []*Node
	count int
}

// NOTE: creates a linked list and returns the interface/contract
func NewLinkedList() LinkedListInterface {
	return &LinkedList{}
}
