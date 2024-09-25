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
			ExpectedNode: node.MakeMapNode(
				yaml.StyleNode(node.MakeStringNode("emptyKey"), 0),
				node.EmptyNode{},

				yaml.StyleNode(node.MakeStringNode("nullKey"), 0),
				node.EmptyNode{},

				yaml.StyleNode(node.MakeStringNode("singleCharNullKey"), 0),
				node.EmptyNode{},

				yaml.StyleNode(node.MakeStringNode("singleQuotedEmptyKey"), 0),
				yaml.StyleNode(
					node.MakeStringNode(""),
					pkgyaml.SingleQuotedStyle,
				),

				yaml.StyleNode(node.MakeStringNode("singleQuotedNullKey"), 0),
				yaml.StyleNode(
					node.MakeStringNode("null"),
					pkgyaml.SingleQuotedStyle,
				),

				yaml.StyleNode(
					node.MakeStringNode("singleQuotedSingleCharNullKey"),
					0,
				),

				yaml.StyleNode(
					node.MakeStringNode("~"),
					pkgyaml.SingleQuotedStyle,
				),

				yaml.StyleNode(
					node.MakeStringNode("doubleQuotedEmptyKey"),
					0,
				),
				yaml.StyleNode(
					node.MakeStringNode(""),
					pkgyaml.DoubleQuotedStyle,
				),

				yaml.StyleNode(
					node.MakeStringNode("doubleQuotedNullKey"),
					0,
				),
				yaml.StyleNode(
					node.MakeStringNode("null"),
					pkgyaml.DoubleQuotedStyle,
				),

				yaml.StyleNode(
					node.MakeStringNode("doubleQuotedSingleCharNullKey"),
					0,
				),
				yaml.StyleNode(
					node.MakeStringNode("~"),
					pkgyaml.DoubleQuotedStyle,
				),
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
			ExpectedNode: node.MakeMapNode(
				yaml.StyleNode(node.MakeStringNode("trueValue"), 0),
				node.MakeBoolNode(true),
				yaml.StyleNode(node.MakeStringNode("falseValue"), 0),
				node.MakeBoolNode(false),
				yaml.StyleNode(node.MakeStringNode("integerValue"), 0),
				node.MakeIntegerNode(1008001),
				yaml.StyleNode(node.MakeStringNode("floatValue"), 0),
				node.MakeFloatNode(3.1),
				yaml.StyleNode(node.MakeStringNode("stringValue"), 0),
				yaml.StyleNode(node.MakeStringNode("aboba"), 0),
			),
		},
	)
}
