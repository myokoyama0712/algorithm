package tree

type Node struct {
	id       int
	children []Node
}

func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.children = []Node{}
	return node
}

func (n *Node) toString() string {
	return "id: 1, children: []"
}
