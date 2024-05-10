package number_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Decode_Number(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:     "1234 integer",
			Data:         `1234`,
			ExpectedNode: node.MakeIntegerNode(1234),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "1234 integer in array",
			Data:         `[1234]`,
			ExpectedNode: node.MakeArrayNodeWithContent(node.MakeIntegerNode(1234)),
		},
		&jsontest.DecodeTestCase{
			CaseName: "integer array",
			Data:     `[1234,4321,5678,8765]`,
			ExpectedNode: node.MakeArrayNodeWithContent(
				node.MakeIntegerNode(1234),
				node.MakeIntegerNode(4321),
				node.MakeIntegerNode(5678),
				node.MakeIntegerNode(8765),
			),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "negative integer",
			Data:         "-2138315144744848",
			ExpectedNode: node.MakeIntegerNode(-2138315144744848),
		},
		&jsontest.DecodeTestCase{
			CaseName: "negative integer array",
			Data:     "[-2138315144744848, -1238313, 123921321, 43934914]",
			ExpectedNode: node.MakeArrayNodeWithContent(
				node.MakeIntegerNode(-2138315144744848),
				node.MakeIntegerNode(-1238313),
				node.MakeIntegerNode(123921321),
				node.MakeIntegerNode(43934914),
			),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "base float",
			Data:         "1234.5",
			ExpectedNode: node.MakeFloatNode(1234.5),
		},
		&jsontest.DecodeTestCase{
			CaseName: "base float array",
			Data:     "[1.123123,123123.123412]",
			ExpectedNode: node.MakeArrayNodeWithContent(
				node.MakeFloatNode(1.123123),
				node.MakeFloatNode(123123.123412),
			),
		},
	)
}
