package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
)

type MapSetTest struct {
	CaseName        string
	MapNode         node.MapNode
	Key, Value      node.Node
	ExpectedMapNode node.MapNode
}

func (m *MapSetTest) Name() string {
	return m.CaseName
}

func (m *MapSetTest) Test(t *testing.T) {
	actualMapNode := node.MapSet(m.MapNode, m.Key, m.Value)

	assert.Equal(t, m.ExpectedMapNode, actualMapNode)
}

func Test_MapSet(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapSetTest{
			CaseName: "empty map test",
			MapNode:  node.MakeMapNode(),
			Key:      node.MakeStringNode("Key"),
			Value:    node.MakeStringNode("Value"),
			ExpectedMapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("Value"),
			),
		},
		&MapSetTest{
			CaseName: "map with exists key",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("Value"),
			),
			Key:   node.MakeStringNode("Key"),
			Value: node.MakeStringNode("NewValue"),
			ExpectedMapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("NewValue"),
			),
		},
		&MapSetTest{
			CaseName: "map with key without value",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
			),
			Key:   node.MakeStringNode("Key"),
			Value: node.MakeStringNode("NewValue"),
			ExpectedMapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("NewValue"),
			),
		},
		&MapSetTest{
			CaseName: "many keys map with target key in the middle",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("akeka"),
				node.MakeStringNode("value"),
				node.MakeStringNode("akeka"),
			),
			Key:   node.MakeStringNode("aboba"),
			Value: node.MakeStringNode("aboba"),
			ExpectedMapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("value"),
				node.MakeStringNode("akeka"),
			),
		},
		&MapSetTest{
			CaseName: "many keys map with target key in the end",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("akeka"),
				node.MakeStringNode("value"),
				node.MakeStringNode("akeka"),
			),
			Key:   node.MakeStringNode("value"),
			Value: node.MakeStringNode("aboba"),
			ExpectedMapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("akeka"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
			),
		},
	)
}

type MapSearchByStringKeyTest struct {
	CaseName      string
	MapNode       node.MapNode
	SearchKey     string
	ExpectedIndex int
}

func (m *MapSearchByStringKeyTest) Name() string {
	return m.CaseName
}

func (m *MapSearchByStringKeyTest) Test(t *testing.T) {
	index := node.MapSearchByStringKey(m.MapNode, m.SearchKey)

	assert.Equal(t, m.ExpectedIndex, index)
}

func Test_MapSearchByStringKey(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapSearchByStringKeyTest{
			CaseName:      "empty map node",
			MapNode:       node.MakeMapNode(),
			SearchKey:     "any",
			ExpectedIndex: -1,
		},
		&MapSearchByStringKeyTest{
			CaseName: "not exists key",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("hehe"),
				node.MakeIntegerNode(1),
			),
			SearchKey:     "any",
			ExpectedIndex: -1,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("hehe"),
				node.MakeIntegerNode(1),
			),
			SearchKey:     "aboba",
			ExpectedIndex: 3,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("true"),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
			),
			SearchKey:     "hehepopa",
			ExpectedIndex: 7,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("true"),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
			),
			SearchKey:     "true",
			ExpectedIndex: 5,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key by wrong kind",
			MapNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeBoolNode(true),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
			),
			SearchKey:     "true",
			ExpectedIndex: -1,
		},
	)
}
