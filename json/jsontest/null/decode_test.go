package null_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

//go:embed testdata/valid/valid_object.json
var validNull string

func Test_Null_Decode_Success(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName: "valid",
			Data:     validNull,
			ExpectedNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("null"),
				node.EmptyNode{},

				node.MakeStringNode("array"),
				node.MakeArrayNodeWithContent(
					node.EmptyNode{},
					node.EmptyNode{},
					node.EmptyNode{},
					node.EmptyNode{},
					node.MakeMapNodeWithContent(
						node.MakeStringNode("key"),
						node.EmptyNode{},
					),
				),
				node.MakeStringNode("object"),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("value"),
					node.EmptyNode{},
				),
			),
		},
	)
}

func Test_Null_Decode_Failure(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:    "n",
			Data:        `[n]`,
			ExpectedErr: json.ErrNullNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "nn",
			Data:        `[nn]`,
			ExpectedErr: json.ErrNullNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "nnn",
			Data:        `[nnn]`,
			ExpectedErr: json.ErrNullNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "nnnn",
			Data:        `[nnnn]`,
			ExpectedErr: json.ErrNullNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:     "null",
			Data:         `[null]`,
			ExpectedNode: node.MakeArrayNodeWithContent(node.EmptyNode{}),
		},
	)
}
