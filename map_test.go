package node_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/require"
)

type MapAppendTest struct {
	CaseName        string
	InitNodes       []node.Node
	Key             node.Node
	Value           node.Node
	ExpectedContent []node.Node
}

func (m *MapAppendTest) Name() string {
	return m.CaseName
}

func (m *MapAppendTest) Test(t *testing.T) {
	mapNode := node.MakeMapNode(m.InitNodes...)

	mapNode = node.MapAppend(mapNode, m.Key, m.Value)

	require.Equal(t, m.ExpectedContent, mapNode.Content())
}

func Test_MapAppend(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapAppendTest{
			CaseName:        "zero init nodes",
			InitNodes:       []node.Node{},
			Key:             node.MakeStringNode("key"),
			Value:           node.MakeStringNode("value"),
			ExpectedContent: []node.Node{node.MakeStringNode("key"), node.MakeStringNode("value")},
		},
		&MapAppendTest{
			CaseName:  "single init node",
			InitNodes: []node.Node{node.MakeStringNode("empty")},
			Key:       node.MakeStringNode("key"),
			Value:     node.MakeStringNode("value"),
			ExpectedContent: []node.Node{
				node.MakeStringNode("empty"),
				node.EmptyNode{},

				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			},
		},
		&MapAppendTest{
			CaseName: "two init nodes",
			InitNodes: []node.Node{
				node.MakeStringNode("not_empty"),
				node.MakeStringNode("null"),
			},
			Key:   node.MakeStringNode("key"),
			Value: node.MakeStringNode("value"),
			ExpectedContent: []node.Node{
				node.MakeStringNode("not_empty"),
				node.MakeStringNode("null"),

				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			},
		},
	)
}
