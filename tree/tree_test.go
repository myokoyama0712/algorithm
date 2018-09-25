package tree

import (
	"fmt"
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
	node := NewNode(1)
	node.AddChild(2)
	node.AddChild(3)
	node.children[0].AddChild(4)
	node.children[0].AddChild(5)
	node.children[1].AddChild(6)
	node.children[1].AddChild(7)
	node.children[1].AddChild(8)
	expectedString := "id: 1, children: [2, 3]\n"
	expectedString += "\tid: 2, children: [4, 5]\n"
	expectedString += "\t\tid: 4, children: []\n"
	expectedString += "\t\tid: 5, children: []\n"
	expectedString += "\tid: 3, children: [6, 7, 8]\n"
	expectedString += "\t\tid: 6, children: []\n"
	expectedString += "\t\tid: 7, children: []\n"
	expectedString += "\t\tid: 8, children: []\n"
	fmt.Println(expectedString)
	assert.Equal(t, expectedString, node.ToString())
}
