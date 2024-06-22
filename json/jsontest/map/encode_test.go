package map_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Map_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "empty map",
			Node:         node.MakeMapNode(),
			ExpectedData: "{}",
		},
		&jsontest.EncodeTestCase{
			CaseName: "basic map",
			Node: node.MakeMapNode(
				node.MakeStringNode("key"), node.MakeStringNode("value"),
			),
			ExpectedData: `{"key":"value"}`,
		},
		&jsontest.EncodeTestCase{
			CaseName: "inner map",
			Node: node.MakeMapNode(
				node.MakeStringNode("key"), node.MakeStringNode("value"),
				node.MakeStringNode("object"), node.MakeMapNode(
					node.MakeStringNode("key"), node.MakeStringNode("value"),
					node.MakeStringNode("object"), node.MakeMapNode(),
				),
				node.MakeStringNode("empty"), node.MakeMapNode(),
			),
			ExpectedData: `{"key":"value","object":{"key":"value","object":{}},"empty":{}}`,
		},
		&jsontest.EncodeTestCase{
			CaseName: "bool key kind",
			Node: node.MakeMapNode(
				node.MakeBoolNode(false), node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
		&jsontest.EncodeTestCase{
			CaseName: "null key kind",
			Node: node.MakeMapNode(
				node.EmptyNode{}, node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
		&jsontest.EncodeTestCase{
			CaseName: "array key kind",
			Node: node.MakeMapNode(
				node.MakeArrayNode(), node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
		&jsontest.EncodeTestCase{
			CaseName: "map key kind",
			Node: node.MakeMapNode(
				node.MakeMapNode(), node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
		&jsontest.EncodeTestCase{
			CaseName: "integer key kind",
			Node: node.MakeMapNode(
				node.MakeIntegerNode(0), node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
		&jsontest.EncodeTestCase{
			CaseName: "float key kind",
			Node: node.MakeMapNode(
				node.MakeFloatNode(0), node.MakeStringNode("ksdf"),
			),
			ExpectedErr: json.ErrInvalidMapKeyKind,
		},
	)
}
