package array_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Array_Decode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:    "non closed array",
			Data:        "[[]",
			ExpectedErr: json.ErrArrayNotClosed,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "single dot in array",
			Data:        "[.]",
			ExpectedErr: json.NewErrUnexpectedByte('.'),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "empty array",
			Data:         "[]",
			ExpectedNode: node.MakeArrayNode(),
		},
		&jsontest.DecodeTestCase{
			CaseName: "inner array",
			Data:     "[[],[],[[],[]]]",
			ExpectedNode: node.MakeArrayNodeWithContent(
				node.MakeArrayNode(),
				node.MakeArrayNode(),
				node.MakeArrayNodeWithContent(
					node.MakeArrayNode(),
					node.MakeArrayNode(),
				),
			),
		},
		&jsontest.DecodeTestCase{
			CaseName: "array with all types of node",
			Data:     `[null,true,false,1231323,123.1,"Hello World!",[],{}]`,
			ExpectedNode: node.MakeArrayNodeWithContent(
				node.EmptyNode{},
				node.MakeBoolNode(true),
				node.MakeBoolNode(false),
				node.MakeIntegerNode(1231323),
				node.MakeFloatNode(123.1),
				node.MakeStringNode("Hello World!"),
				node.MakeArrayNode(),
				node.MakeMapNode(),
			),
		},
	)
}
