package map_test

import (
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Map_Decode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:    "non closed array",
			Data:        `{"hello":{}`,
			ExpectedErr: json.ErrMapNotClosed,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "dot in map value",
			Data:        `{"hello":.}`,
			ExpectedErr: json.NewErrUnexpectedByte('.'),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "',' in map key",
			Data:        "{,}",
			ExpectedErr: json.NewErrUnexpectedByte(','),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "dot in map key",
			Data:        "{.}",
			ExpectedErr: json.NewErrUnexpectedByte('.'),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "empty map",
			Data:         "{}",
			ExpectedNode: node.MakeMapNode(),
		},
		&jsontest.DecodeTestCase{
			CaseName: "inner empty map",
			Data:     `{"a":{"b":{"c":{}}},"b":{}}`,
			ExpectedNode: node.MakeMapNodeWithContent(
				node.MakeStringNode("a"),
				node.MakeMapNodeWithContent(
					node.MakeStringNode("b"),
					node.MakeMapNodeWithContent(
						node.MakeStringNode("c"),
						node.MakeMapNode(),
					),
				),
				node.MakeStringNode("b"),
				node.MakeMapNode(),
			),
		},
	)
}
