package test_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/yaml"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pkgyaml "gopkg.in/yaml.v3"

	_ "embed"
)

var (
	//go:embed testdata/null.yaml
	nullData []byte
	//go:embed testdata/value_kind.yaml
	valueData []byte
)

type DecodeTester struct {
	CaseName     string
	Decoder      node.Decoder
	Input        []byte
	ExpectedNode node.Node
}

func (d *DecodeTester) Name() string {
	return d.CaseName
}

func (d *DecodeTester) Test(t *testing.T) {
	node, err := d.Decoder.Decode(d.Input)
	require.NoError(t, err)

	assert.Equal(t, d.ExpectedNode, node)
}

func Test_DecodeEmpty(t *testing.T) {
	tester.RunNamedTesters(t,
		&DecodeTester{
			CaseName: "all null yaml cases",
			Decoder:  new(yaml.Decoder),
			Input:    nullData,
			ExpectedNode: node.MakeMapNodeWithContent(
				&yaml.StyleNode{
					Node: node.MakeStringNode("emptyKey"),
				},
				node.EmptyNode{},

				&yaml.StyleNode{
					Node: node.MakeStringNode("nullKey"),
				},
				node.EmptyNode{},

				&yaml.StyleNode{
					Node: node.MakeStringNode("singleCharNullKey"),
				},
				node.EmptyNode{},

				&yaml.StyleNode{
					Node: node.MakeStringNode("singleQuotedEmptyKey"),
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode(""),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.StyleNode{
					Node: node.MakeStringNode("singleQuotedNullKey"),
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode("null"),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.StyleNode{
					Node: node.MakeStringNode("singleQuotedSingleCharNullKey"),
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode("~"),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.StyleNode{
					Node: node.MakeStringNode("doubleQuotedEmptyKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode(""),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},

				&yaml.StyleNode{
					Node: node.MakeStringNode("doubleQuotedNullKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode("null"),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},

				&yaml.StyleNode{
					Node: node.MakeStringNode("doubleQuotedSingleCharNullKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.StyleNode{
					Node:      node.MakeStringNode("~"),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},
			),
		},
	)
}

func Test_DecodeValueKind(t *testing.T) {
	tester.RunNamedTesters(t,
		&DecodeTester{
			CaseName: "bool, integer, float value kinds",
			Decoder:  new(yaml.Decoder),
			Input:    valueData,
			ExpectedNode: node.MakeMapNodeWithContent(
				&yaml.StyleNode{Node: node.MakeStringNode("trueValue")},
				node.MakeBoolNode(true),
				&yaml.StyleNode{Node: node.MakeStringNode("falseValue")},
				node.MakeBoolNode(false),
				&yaml.StyleNode{Node: node.MakeStringNode("integerValue")},
				node.MakeIntegerNode(1008001),
				&yaml.StyleNode{Node: node.MakeStringNode("floatValue")},
				node.MakeFloatNode(3.1),
				&yaml.StyleNode{Node: node.MakeStringNode("stringValue")},
				&yaml.StyleNode{Node: node.MakeStringNode("aboba")},
			),
		},
	)
}
