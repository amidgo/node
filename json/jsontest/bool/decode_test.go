package bool_test

import (
	_ "embed"
	"testing"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Decode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.DecodeTestCase{
			CaseName:     "true",
			Data:         "true",
			ExpectedNode: node.MakeBoolNode(true),
		},
		&jsontest.DecodeTestCase{
			CaseName:     "false",
			Data:         "false",
			ExpectedNode: node.MakeBoolNode(false),
		},
		&jsontest.DecodeTestCase{
			CaseName:    "fale",
			Data:        "fale",
			ExpectedErr: json.ErrFalseNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "trea",
			Data:        "trea",
			ExpectedErr: json.ErrTrueNotValid,
		},
	)
}
