package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			ExpectedMapNode: node.MakeMapNode(
				node.MakeStringNode("Key"),
				node.MakeStringNode("Value"),
			),
		},
		&MapSetTest{
			CaseName: "map with exists key",
			MapNode: node.MakeMapNode(
				node.MakeStringNode("Key"),
				node.MakeStringNode("Value"),
			),
			Key:   node.MakeStringNode("Key"),
			Value: node.MakeStringNode("NewValue"),
			ExpectedMapNode: node.MakeMapNode(
				node.MakeStringNode("Key"),
				node.MakeStringNode("NewValue"),
			),
		},
		&MapSetTest{
			CaseName: "many keys map with target key in the middle",
			MapNode: node.MakeMapNode(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("akeka"),
				node.MakeStringNode("value"),
				node.MakeStringNode("akeka"),
			),
			Key:   node.MakeStringNode("aboba"),
			Value: node.MakeStringNode("aboba"),
			ExpectedMapNode: node.MakeMapNode(
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
			MapNode: node.MakeMapNode(
				node.MakeStringNode("Key"),
				node.MakeStringNode("value"),
				node.MakeStringNode("aboba"),
				node.MakeStringNode("akeka"),
				node.MakeStringNode("value"),
				node.MakeStringNode("akeka"),
			),
			Key:   node.MakeStringNode("value"),
			Value: node.MakeStringNode("aboba"),
			ExpectedMapNode: node.MakeMapNode(
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
			MapNode: node.MakeMapNode(
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
			MapNode: node.MakeMapNode(
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
			MapNode: node.MakeMapNode(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("true"),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
				node.MakeStringNode("hehepopavalue"),
			),
			SearchKey:     "hehepopa",
			ExpectedIndex: 7,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key",
			MapNode: node.MakeMapNode(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeStringNode("true"),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
				node.MakeStringNode("hehepopavalue"),
			),
			SearchKey:     "true",
			ExpectedIndex: 5,
		},
		&MapSearchByStringKeyTest{
			CaseName: "exists key by wrong kind",
			MapNode: node.MakeMapNode(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),
				node.MakeStringNode("aboba"),
				node.MakeBoolNode(true),
				node.MakeBoolNode(true),
				node.MakeIntegerNode(1),
				node.MakeStringNode("hehepopa"),
				node.MakeStringNode("hehepopavalue"),
			),
			SearchKey:     "true",
			ExpectedIndex: -1,
		},
	)
}

type MapIterableGenerateTest struct {
	CaseName      string
	SourceMap     node.MapNode
	IterationStep node.IterationStep
	ExpectedMap   node.MapNode
	ExpectedError error
}

func (m *MapIterableGenerateTest) Name() string {
	return m.CaseName
}

func (m *MapIterableGenerateTest) Test(t *testing.T) {
	gen := node.NewMapIterableGenerate(m.SourceMap, m.IterationStep)

	mapNode, err := gen.MapNode()
	require.ErrorIs(t, err, m.ExpectedError)
	assert.Equal(t, m.ExpectedMap, mapNode)
}

type IterationSetStep struct {
	Key                string
	KeyNode, ValueNode node.Node
}

func (i *IterationSetStep) KeyValue(key, value node.Node) (resKey, resValue node.Node, err error) {
	if key.Value() == i.Key {
		return i.KeyNode, i.ValueNode, nil
	}

	return key, value, nil
}

func Test_MapIterableGenerate(t *testing.T) {
	tester.RunNamedTesters(t,
		&MapIterableGenerateTest{
			CaseName: "zero steps",
			SourceMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			),
			IterationStep: node.NewJoinIterationStep(),
			ExpectedMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			),
		},
		&MapIterableGenerateTest{
			CaseName: "single step",
			SourceMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			),
			IterationStep: node.NewJoinIterationStep(
				&IterationSetStep{
					Key:       "key",
					KeyNode:   node.MakeStringNode("key"),
					ValueNode: node.MakeStringNode("newValue"),
				},
			),
			ExpectedMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("newValue"),
			),
		},
		&MapIterableGenerateTest{
			CaseName: "many steps that handle the same key",
			SourceMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			),
			IterationStep: node.NewJoinIterationStep(
				&IterationSetStep{
					Key:       "key",
					KeyNode:   node.MakeStringNode("key"),
					ValueNode: node.MakeStringNode("newValue"),
				},
				&IterationSetStep{
					Key:       "key",
					KeyNode:   node.MakeStringNode("key"),
					ValueNode: node.MakeStringNode("radicallyNewValue"),
				},
				&IterationSetStep{
					Key:       "aboba",
					KeyNode:   node.MakeStringNode("key"),
					ValueNode: node.MakeStringNode("newValue"),
				},
			),
			ExpectedMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("radicallyNewValue"),
			),
		},
		&MapIterableGenerateTest{
			CaseName: "step that changes key",
			SourceMap: node.MakeMapNode(
				node.MakeStringNode("key"),
				node.MakeStringNode("value"),
			),
			IterationStep: node.NewJoinIterationStep(
				&IterationSetStep{
					Key:       "key",
					KeyNode:   node.MakeStringNode("newKey"),
					ValueNode: node.MakeStringNode("newValue"),
				},
				&IterationSetStep{
					Key:       "newKey",
					KeyNode:   node.MakeStringNode("returnToOldKey"),
					ValueNode: node.MakeStringNode("newValue"),
				},
				&IterationSetStep{
					Key:       "aboba",
					KeyNode:   node.MakeStringNode("key"),
					ValueNode: node.MakeStringNode("newValue"),
				},
			),
			ExpectedMap: node.MakeMapNode(
				node.MakeStringNode("returnToOldKey"),
				node.MakeStringNode("newValue"),
			),
		},
	)
}
