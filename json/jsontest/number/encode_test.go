package number_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Number_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "1234 integer",
			ExpectedData: `1234`,
			Node:         node.MakeIntegerNode(1234),
		},
		&jsontest.EncodeTestCase{
			CaseName:     "1234 integer in array",
			ExpectedData: `[1234]`,
			Node:         node.MakeArrayNodeWithContent(node.MakeIntegerNode(1234)),
		},
		&jsontest.EncodeTestCase{
			CaseName:     "integer array",
			ExpectedData: `[1234,4321,5678,8765]`,
			Node: node.MakeArrayNodeWithContent(
				node.MakeIntegerNode(1234),
				node.MakeIntegerNode(4321),
				node.MakeIntegerNode(5678),
				node.MakeIntegerNode(8765),
			),
		},
		&jsontest.EncodeTestCase{
			CaseName:     "negative integer",
			ExpectedData: "-2138315144744848",
			Node:         node.MakeIntegerNode(-2138315144744848),
		},
		&jsontest.EncodeTestCase{
			CaseName:     "negative integer array",
			ExpectedData: "[-2138315144744848,-1238313,123921321,43934914]",
			Node: node.MakeArrayNodeWithContent(
				node.MakeIntegerNode(-2138315144744848),
				node.MakeIntegerNode(-1238313),
				node.MakeIntegerNode(123921321),
				node.MakeIntegerNode(43934914),
			),
		},
		&jsontest.EncodeTestCase{
			CaseName:    "invalid integer node",
			Node:        &node.UnsafeNode{NType: node.Value, NKind: node.Integer},
			ExpectedErr: json.ErrInvalidIntegerNode,
		},
		&jsontest.EncodeTestCase{
			CaseName:     "base float",
			ExpectedData: "1234.5",
			Node:         node.MakeFloatNode(1234.5),
		},
		&jsontest.EncodeTestCase{
			CaseName:     "base float array",
			ExpectedData: "[1.123123,123123.123412]",
			Node: node.MakeArrayNodeWithContent(
				node.MakeFloatNode(1.123123),
				node.MakeFloatNode(123123.123412),
			),
		},
		&jsontest.EncodeTestCase{
			CaseName:    "invalid float node",
			Node:        &node.UnsafeNode{NType: node.Value, NKind: node.Float},
			ExpectedErr: json.ErrInvalidFloatNode,
		},
	)
}
