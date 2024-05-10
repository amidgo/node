package bool_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

//go:embed testdata/valid/valid_object_flat.json
var flatValidObjectBool string

func Test_Bool_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "false",
			Node:         node.MakeBoolNode(false),
			ExpectedData: "false",
		},
		&jsontest.EncodeTestCase{
			CaseName:     "true",
			Node:         node.MakeBoolNode(true),
			ExpectedData: "true",
		},
		&jsontest.EncodeTestCase{
			CaseName:    "not valid bool node",
			Node:        &node.UnsafeNode{NType: node.Value, NKind: node.Bool},
			ExpectedErr: json.ErrInvalidBoolNode,
		},
		&jsontest.EncodeTestCase{
			CaseName: "strong object with bool",
			Node: node.MakeMapNodeWithContent(
				node.MakeStringNode("hello"),
				node.MakeBoolNode(false),

				node.MakeStringNode("Hello"),
				node.MakeBoolNode(true),

				node.MakeStringNode("object"),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("true"),
					node.MakeBoolNode(true),

					node.MakeStringNode("false"),
					node.MakeBoolNode(false),

					node.MakeStringNode("array"),
					node.MakeArrayNodeWithContent(
						node.MakeBoolNode(false),
					),
				),
				node.MakeStringNode("array"),
				node.MakeArrayNodeWithContent(
					node.MakeBoolNode(true),
					node.MakeBoolNode(false),
					node.MakeBoolNode(true),
					node.MakeBoolNode(true),
					node.MakeBoolNode(false),

					node.MakeMapNodeWithContent(
						node.MakeStringNode("value"),
						node.MakeBoolNode(false),
					),
				),
			),
			ExpectedData: flatValidObjectBool,
		},
	)
}
