package bool_test

import (
	"testing"

	_ "embed"

	"github.com/amidgo/node"
	"github.com/amidgo/node/json"
	"github.com/amidgo/node/json/jsontest"
	"github.com/amidgo/tester"
)

type BoolNode struct{}

func (i BoolNode) Value() string {
	return ""
}

func (i BoolNode) Content() []node.Node {
	return nil
}

func (i BoolNode) Kind() node.Kind {
	return node.Bool
}

func (i BoolNode) Type() node.Type {
	return node.Value
}

func Test_Bool_Encode(t *testing.T) {
	tester.RunNamedTesters(t,
		&jsontest.EncodeTestCase{
			CaseName:     "true",
			Node:         node.MakeBoolNode(true),
			ExpectedData: "true",
		},
		&jsontest.EncodeTestCase{
			CaseName:     "false",
			ExpectedData: "false",
			Node:         node.MakeBoolNode(false),
		},
		&jsontest.EncodeTestCase{
			CaseName:    "invalid bool node",
			Node:        BoolNode{},
			ExpectedErr: json.ErrInvalidBoolNode,
		},
	)
}
