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
