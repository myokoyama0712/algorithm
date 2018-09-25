package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputNodeStruct(t *testing.T) {
	id := 1
	node := NewNode(id)
	assert.Equal(t, "id: 1, children: []", node.toString())
}
