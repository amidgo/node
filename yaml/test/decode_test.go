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
				node.MakeStringNode("emptyKey"),
				node.EmptyNode{},

				node.MakeStringNode("nullKey"),
				node.EmptyNode{},

				node.MakeStringNode("singleCharNullKey"),
				node.EmptyNode{},

				node.MakeStringNode("singleQuotedEmptyKey"),
				yaml.StyleNode(
					node.MakeStringNode(""),
					pkgyaml.SingleQuotedStyle,
				),

				node.MakeStringNode("singleQuotedNullKey"),
				yaml.StyleNode(
					node.MakeStringNode("null"),
					pkgyaml.SingleQuotedStyle,
				),

				node.MakeStringNode("singleQuotedSingleCharNullKey"),

				yaml.StyleNode(
					node.MakeStringNode("~"),
					pkgyaml.SingleQuotedStyle,
				),

				node.MakeStringNode("doubleQuotedEmptyKey"),
				yaml.StyleNode(
					node.MakeStringNode(""),
					pkgyaml.DoubleQuotedStyle,
				),

				node.MakeStringNode("doubleQuotedNullKey"),
				yaml.StyleNode(
					node.MakeStringNode("null"),
					pkgyaml.DoubleQuotedStyle,
				),

				node.MakeStringNode("doubleQuotedSingleCharNullKey"),
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
				node.MakeStringNode("trueValue"),
				node.MakeBoolNode(true),
				node.MakeStringNode("falseValue"),
				node.MakeBoolNode(false),
				node.MakeStringNode("integerValue"),
				node.MakeIntegerNode(1008001),
				node.MakeStringNode("floatValue"),
				node.MakeFloatNode(3.1),
				node.MakeStringNode("stringValue"),
				node.MakeStringNode("aboba"),
			),
		},
	)
}
