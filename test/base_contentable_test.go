package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/stretchr/testify/assert"
)

func TestSetContent(t *testing.T) {
	contentableNode := new(node.BaseContentableNode)

	content := []node.Node{
		&node.UnsafeNode{
			NValue: "dafsk",
		},
		&node.UnsafeNode{
			NValue: "dafsdsk",
		},
	}

	contentableNode.SetContent(content)

	assert.Equal(t, content, contentableNode.Content())
}

func TestAppendNode(t *testing.T) {
	contentableNode := new(node.BaseContentableNode)

	content := []node.Node{
		&node.UnsafeNode{
			NValue: "dafsk",
		},
		&node.UnsafeNode{
			NValue: "dafsdsk",
		},
	}

	for i := range content {
		contentableNode.AppendNode(content[i])
	}

	assert.Equal(t, content, contentableNode.Content())
}
