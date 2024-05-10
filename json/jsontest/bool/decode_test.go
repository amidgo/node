package bool_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

//go:embed testdata/valid/valid_object.json
var validObjectBool string

func Test_Bool_Valid_Object(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName: "strong object case",
			Data:     validObjectBool,
			ExpectedNode: node.MakeMapNodeWithContent(
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
		},
	)
}

func Test_Bool_Decode_Failure(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:    "t value",
			Data:        `{"value":t}`,
			ExpectedErr: json.ErrTrueNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "tt value",
			Data:        `{"value":tt}`,
			ExpectedErr: json.ErrTrueNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "ttt value",
			Data:        `{"value":ttt}`,
			ExpectedErr: json.ErrTrueNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "tttt value",
			Data:        `{"value":tttt}`,
			ExpectedErr: json.ErrTrueNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "ttttt value",
			Data:        `{"value":ttttt}`,
			ExpectedErr: json.ErrTrueNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:     "valid value",
			Data:         `{"value":true}`,
			ExpectedNode: node.MakeMapNodeWithContent(node.MakeStringNode("value"), node.MakeBoolNode(true)),
		},

		&jsontest.DecodeTestCase{
			CaseName:    "f value",
			Data:        `{"value":f}`,
			ExpectedErr: json.ErrFalseNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "ff value",
			Data:        `{"value":ff}`,
			ExpectedErr: json.ErrFalseNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "fff value",
			Data:        `{"value":fff}`,
			ExpectedErr: json.ErrFalseNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "ffff value",
			Data:        `{"value":ffff}`,
			ExpectedErr: json.ErrFalseNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "ttttt value",
			Data:        `{"value":fffff}`,
			ExpectedErr: json.ErrFalseNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:     "false value",
			Data:         `{"value":false}`,
			ExpectedNode: node.MakeMapNodeWithContent(node.MakeStringNode("value"), node.MakeBoolNode(false)),
		},
	)
}
