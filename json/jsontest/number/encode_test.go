package number_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

type IntegerNode struct{}

func (i IntegerNode) Value() string {
	return ""
}

func (i IntegerNode) Content() []node.Node {
	return nil
}

func (i IntegerNode) Kind() node.Kind {
	return node.Integer
}

func (i IntegerNode) Type() node.Type {
	return node.Value
}

type FloatNode struct{}

func (i FloatNode) Value() string {
	return ""
}

func (i FloatNode) Content() []node.Node {
	return nil
}

func (i FloatNode) Kind() node.Kind {
	return node.Float
}

func (i FloatNode) Type() node.Type {
	return node.Value
}

func Test_Number_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "1234 integer",
			Node:         node.MakeIntegerNode(1234),
			ExpectedData: `1234`,
		},
		&jsontest.EncodeTestCase{
			CaseName:     "negative integer",
			Node:         node.MakeIntegerNode(-2138315144744848),
			ExpectedData: "-2138315144744848",
		},
		&jsontest.EncodeTestCase{
			CaseName:     "1234.5 float",
			Node:         node.MakeFloatNode(1234.5),
			ExpectedData: "1234.5",
		},
		&jsontest.EncodeTestCase{
			CaseName:     "-1234.5 float",
			Node:         node.MakeFloatNode(-1234.5),
			ExpectedData: "-1234.5",
		},
		&jsontest.EncodeTestCase{
			CaseName:    "invalid integer node",
			Node:        IntegerNode{},
			ExpectedErr: json.ErrInvalidIntegerNode,
		},
		&jsontest.EncodeTestCase{
			CaseName:    "invalid float node",
			Node:        FloatNode{},
			ExpectedErr: json.ErrInvalidFloatNode,
		},
	)
}
