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
			ExpectedErr: json.ErrContentableNodeNotClosed,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "single dot in array",
			Data:        "[.]",
			ExpectedErr: json.ErrUnexpectedByte,
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
	)
}
