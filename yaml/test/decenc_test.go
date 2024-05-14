package test_test

import (
	_ "embed"
	"fmt"
	"testing"

	pkgyaml "gopkg.in/yaml.v3"

	"github.com/amidgo/node"
	"github.com/amidgo/node/yaml"
	"github.com/amidgo/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/swagger.yaml
	swaggerData []byte
	//go:embed testdata/simple.yaml
	simpleSwaggerData []byte
	//go:embed testdata/simple_with_array.yaml
	simpleSwaggerWithArrayData []byte
	//go:embed testdata/null.yaml
	nullData []byte
)

type DecEncTester struct {
	CaseName string
	Indent   int
	Input    []byte
}

func (de *DecEncTester) Name() string {
	return fmt.Sprintf("decenc tester %s", de.CaseName)
}

func (de *DecEncTester) Test(t *testing.T) {
	nd, err := yaml.Decode(de.Input)
	require.NoError(t, err)

	data, err := yaml.EncodeWithIndent(nd, de.Indent)
	require.NoError(t, err)

	t.Logf("encodedLen:%d inputLen:%d", len(data), len(de.Input))
	t.Logf("encoded:\n%s\ninput:\n%s", string(data), string(de.Input))

	assert.Equal(t, string(de.Input), string(data))
}

func Test_DecEnc(t *testing.T) {
	tester.RunNamedTesters(t,
		&DecEncTester{
			CaseName: "simple swagger 3.0.1",
			Indent:   2,
			Input:    simpleSwaggerData,
		},
		&DecEncTester{
			CaseName: "simple swagger with array 3.0.1",
			Indent:   2,
			Input:    simpleSwaggerWithArrayData,
		},
		&DecEncTester{
			CaseName: "full swagger 3.0.1",
			Indent:   2,
			Input:    swaggerData,
		},
	)

}

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
				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("emptyKey"),
				},
				node.EmptyNode{},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("nullKey"),
				},
				node.EmptyNode{},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("singleCharNullKey"),
				},
				node.EmptyNode{},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("singleQuotedEmptyKey"),
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode(""),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("singleQuotedNullKey"),
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode("null"),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("singleQuotedSingleCharNullKey"),
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode("~"),
					YamlStyle: pkgyaml.SingleQuotedStyle,
				},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("doubleQuotedEmptyKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode(""),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("doubleQuotedNullKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode("null"),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},

				&yaml.YamlStyleNode{
					Node: node.MakeStringNode("doubleQuotedSingleCharNullKey"),
					// YamlStyle: pkgyaml.TaggedStyle,
				},
				&yaml.YamlStyleNode{
					Node:      node.MakeStringNode("~"),
					YamlStyle: pkgyaml.DoubleQuotedStyle,
				},
			),
		},
	)
}
