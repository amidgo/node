package null_test

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
			CaseName:    "n",
			Data:        `n`,
			ExpectedErr: json.ErrNullNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:    "nn",
			Data:        `nn`,
			ExpectedErr: json.ErrNullNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "nnn",
			Data:        `nnn`,
			ExpectedErr: json.ErrNullNotValid,
		},

		&jsontest.DecodeTestCase{
			CaseName:    "nnnn",
			Data:        `nnnn`,
			ExpectedErr: json.ErrNullNotValid,
		},
		&jsontest.DecodeTestCase{
			CaseName:     "null",
			Data:         `null`,
			ExpectedNode: node.EmptyNode{},
		},
	)
}
