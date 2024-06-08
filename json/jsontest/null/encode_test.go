package null_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

func Test_Null_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "null",
			Node:         node.EmptyNode{},
			ExpectedData: `null`,
		},
	)
}
