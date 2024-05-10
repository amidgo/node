package null_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

//go:embed testdata/valid/valid_object_flat.json
var flatValidNull string

func Test_Null_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "null",
			Node:         node.EmptyNode{},
			ExpectedData: "null",
		},
		&jsontest.EncodeTestCase{
			CaseName: "valid object",
			Node: node.MakeMapNodeWithContent(
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
			ExpectedData: flatValidNull,
		},
	)
}
