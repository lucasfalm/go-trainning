package node

type Node struct {
	head  NodeInterface
	tail  NodeInterface
	value any
}

type NodeInterface interface {
	Head() NodeInterface
	SetHead(NodeInterface)
	Tail() NodeInterface
	SetTail(NodeInterface)
	Value() any
	SetValue(any)
	Print()
}

func NewNode() NodeInterface {
	return &Node{}
}
