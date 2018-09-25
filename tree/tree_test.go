package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputEmptyNode(t *testing.T) {
	id := 1
	node := NewNode(id)
	assert.Equal(t, "id: 1, children: []", node.ToString())
}

func TestAddChildNode(t *testing.T) {
	node := NewNode(1)
	node.AddChild(2)
	node.AddChild(3)
	expectedSlice := []*Node{
		&Node{id: 2, children: []*Node{}},
		&Node{id: 3, children: []*Node{}},
	}
	assert.ElementsMatch(t, expectedSlice, node.children)
}

func TestOutputNotEmptyNode(t *testing.T) {
}
