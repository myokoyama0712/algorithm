package tree

import (
	"fmt"
)

/*
Node is struct denoting tree node.
*/
type Node struct {
	id       int
	children []*Node
}

/*
NewNode returns pointer of Node struct.
*/
func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.children = []*Node{}
	return node
}

/*
ToString returns string formatted like
`
id: self, children: [child1, child2, ...]
	id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
	id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
	...
`
*/
func (n *Node) ToString() string {
	return fmt.Sprintf("id: %d, children: %v", n.id, n.children)
}

/*
AddChild append new child node to its childlen.
*/
func (n *Node) AddChild(id int) {
	newChild := new(Node)
	newChild.id = id
	newChild.children = []*Node{}

	n.children = append(n.children, newChild)
}
