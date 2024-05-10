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
			Node: node.MakeArrayNodeWithContent(
				node.MakeArrayNode(),
				node.MakeArrayNode(),
				node.MakeArrayNodeWithContent(
					node.MakeArrayNode(),
					node.MakeArrayNode(),
				),
				node.MakeMapNode(),
			),
			ExpectedData: "[[],[],[[],[]],{}]",
		},
	)
}
