package array_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Array_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "empty array node",
			Node:         node.MakeArrayNode(),
			ExpectedData: "[]",
		},
		&jsontest.EncodeTestCase{
			CaseName: "inner array node",
			Node: node.MakeArrayNode(
				node.MakeArrayNode(),
				node.MakeArrayNode(),
				node.MakeArrayNode(
					node.MakeArrayNode(),
					node.MakeArrayNode(),
				),
				node.MakeMapNode(),
			),
			ExpectedData: "[[],[],[[],[]],{}]",
		},
		&jsontest.EncodeTestCase{
			CaseName: "array with all types of node",
			Node: node.MakeArrayNode(
				node.EmptyNode{},
				node.MakeBoolNode(true),
				node.MakeBoolNode(false),
				node.MakeIntegerNode(1231323),
				node.MakeFloatNode(123.1),
				node.MakeStringNode("Hello World!"),
				node.MakeArrayNode(),
				node.MakeMapNode(),
			),
			ExpectedData: `[null,true,false,1231323,123.1,"Hello World!",[],{}]`,
		},
	)
}
